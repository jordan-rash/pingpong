package main

//go:generate wit-bindgen tiny-go wit/ --out-dir gen/

import (
	"fmt"
	provider "github.com/wasmCloud/provider-sdk-go"
	core "github.com/wasmcloud/interfaces/core/tinygo"
	"time"
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

func handleNewLink(linkdef core.LinkDefinition) error {
	go sendPings(linkdef.ActorId)
	return nil
}

func sendPings(aId string) {
	for {
		b, err := p.ToActor(aId, nil, "PingPong.Ping")
		if err != nil {
			panic(err)
		} else {
			fmt.Println(b)
		}
		time.Sleep(1 * time.Second)
	}
}
