module micro_demo/config-grpc-srv

go 1.13

require (
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/config/source/grpc/v2 v2.9.1
	google.golang.org/grpc v1.30.0
	micro_demo/basic v0.0.0-00010101000000-000000000000
	micro_demo/proto v0.0.0-00010101000000-000000000000
)

replace (
	micro_demo/basic => ../basic
	micro_demo/comm => ../comm
	micro_demo/proto => ../proto
)
