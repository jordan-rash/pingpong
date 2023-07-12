package main

import (
	ping_pong "github.com/jordan-rash/pingpong/gen"
	"github.com/wasmCloud/actor-tinygo"
)

func main() {
	me := Example{}
	actor.RegisterHandlers(ping_pong.PingPongHandler(&me))
}

type Example struct{}

func (e *Example) Ping(ctx *actor.Context) (string, error) {
	return "pong", nil
}
