package ping_pong

type JordanRashPingpongTypesPong = string
type JordanRashPingpongPingpongPong = string

var jordan_rash_pingpong_pingpong ExportsJordanRashPingpongPingpong = nil

func SetExportsJordanRashPingpongPingpong(i ExportsJordanRashPingpongPingpong) {
	jordan_rash_pingpong_pingpong = i
}

type ExportsJordanRashPingpongPingpong interface {
	Ping() string
}
