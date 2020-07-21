module micro_demo/gateway

go 1.14

require (
	github.com/micro-in-cn/starter-kit v1.18.0
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/micro/cors/v2 v2.9.1
	github.com/micro/micro/v2 v2.9.3
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	go.uber.org/zap v1.15.0
	micro_demo/comm v0.0.0-00010101000000-000000000000
)

replace (
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
	micro_demo/basic => ../basic
	micro_demo/comm => ../comm
	micro_demo/proto => ../proto
)
