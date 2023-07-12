package ping_pong

import actor "github.com/wasmCloud/actor-tinygo"

type JordanRashPingpongTypesPong string
type JordanRashPingpongPingpongPong string

type ExportsJordanRashPingpongPingpong interface {
	Ping(ctx *actor.Context) (string, error)
}
