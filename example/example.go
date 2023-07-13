package main

import (
	ping_pong "github.com/jordan-rash/pingpong/gen"
	"github.com/wasmCloud/actor-tinygo"
)

func main() {
	me := Example{
		logger: actor.NewProviderLogging(),
	}

	actor.RegisterHandlers(ping_pong.PingPongHandler(&me))
}

type Example struct {
	logger *actor.LoggingSender
}

func (e *Example) Ping(ctx *actor.Context) (string, error) {
	e.logger.WriteLog(ctx, actor.LogEntry{Level: "info", Text: "im trying to pong"})
	return "pong", nil
}
