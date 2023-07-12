package ping_pong

import (
	actor "github.com/wasmCloud/actor-tinygo"
	msgpack "github.com/wasmcloud/tinygo-msgpack"
)

func (o *JordanRashPingpongTypesPong) MEncode(encoder msgpack.Writer) error {
	encoder.WriteString("pong")
	return encoder.CheckError()
}

func MDecodePong(d *msgpack.Decoder) (JordanRashPingpongTypesPong, error) {
	val, err := d.ReadString()
	if err != nil {
		return "", err
	}

	return JordanRashPingpongTypesPong(val), nil
}

type PingPongSender struct{ transport actor.Transport }
type PingPongReceiver struct{}

func NewProviderPingPong() *PingPongSender {
	transport := actor.ToProvider("jordan-rash:pingpong", "default")
	return &PingPongSender{transport: transport}
}

func PingPongHandler(a ExportsJordanRashPingpongPingpong) actor.Handler {
	return actor.NewHandler("PingPong", &PingPongReceiver{}, a)
}

func (r *PingPongReceiver) Dispatch(ctx *actor.Context, svc interface{}, message *actor.Message) (*actor.Message, error) {
	// svc_, _ := svc.(ExportsJordanRashPingpongPingpong)
	switch message.Method {
	case "Ping":
		{
			var resp JordanRashPingpongTypesPong

			var sizer msgpack.Sizer
			size_enc := &sizer
			resp.MEncode(size_enc)
			buf := make([]byte, sizer.Len())
			encoder := msgpack.NewEncoder(buf)
			enc := &encoder
			resp.MEncode(enc)
			return &actor.Message{Method: "HttpClient.Request", Arg: buf}, nil
		}
	default:
		return nil, actor.NewRpcError("MethodNotHandled", "PingPong."+message.Method)
	}
}
