module github.com/jordan-rash/pingpong

go 1.20

require (
	github.com/wasmCloud/actor-tinygo v0.0.0-00010101000000-000000000000
	github.com/wasmCloud/provider-sdk-go v0.0.0-00010101000000-000000000000
	github.com/wasmcloud/tinygo-msgpack v0.1.4
)

require (
	github.com/bombsimon/logrusr/v3 v3.1.0 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/klauspost/compress v1.16.5 // indirect
	github.com/minio/highwayhash v1.0.2 // indirect
	github.com/nats-io/jwt/v2 v2.4.1 // indirect
	github.com/nats-io/nats.go v1.25.0 // indirect
	github.com/nats-io/nkeys v0.4.4 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/wasmcloud/tinygo-cbor v0.1.0 // indirect
	golang.org/x/crypto v0.8.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/time v0.3.0 // indirect
)

replace github.com/wasmCloud/actor-tinygo => ../actor-tinygo

replace github.com/wasmcloud/actor-tinygo/core => ../actor-tinygo/core

replace github.com/wasmCloud/provider-sdk-go => ../provider-sdk-go
