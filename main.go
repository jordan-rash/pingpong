package main

//go:generate wit-bindgen tiny-go wit/ --out-dir gen/

import (
	"time"

	provider "github.com/wasmCloud/provider-sdk-go"
	wasmcloud_core "github.com/wasmCloud/provider-sdk-go/core"
)

var (
	p *provider.WasmcloudProvider
)

func main() {
	var err error

	p, err = provider.New(
		"jordan-rash:pingpong",
		provider.WithNewLinkFunc(handleNewLink),
		provider.WithHealthCheckMsg(healthCheckMsg),
	)
	if err != nil {
		panic(err)
	}

	err = p.Start()
	if err != nil {
		panic(err)
	}
}

func healthCheckMsg() string {
	return "THE PING PONG provider!"
}

// func handleNewLink(linkdef wasmcloud_core.WasmcloudCoreTypesLinkDefinition) error {
func handleNewLink(linkdef wasmcloud_core.WasmcloudCoreTypesLinkDefinition) error {
	// _, err := p.ToActor(linkdef.ActorId, []byte("ping"), "PingPong.Ping")
	// if err != nil {
	// 	p.Logger.Error(err, "from "+linkdef.ActorId)
	// }
	//
	// return err
	return sendPings(linkdef.ActorId)
}

func sendPings(aId string) error {
	for {
		_, err := p.ToActor(aId, []byte("ping"), "PingPong.Ping")
		if err != nil {
			p.Logger.Error(err, "from "+aId)
		}
		time.Sleep(1 * time.Second)
	}

	return nil
}
