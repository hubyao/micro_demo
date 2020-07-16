module micro_demo/user-srv

go 1.13

replace (
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
	micro_demo/basic => ../basic
	micro_demo/comm => ../comm
	micro_demo/proto => ../proto
)

require (
	github.com/micro/go-micro/v2 v2.9.1
	micro_demo/basic v0.0.0-00010101000000-000000000000
	micro_demo/comm v0.0.0-00010101000000-000000000000
	micro_demo/proto v0.0.0-00010101000000-000000000000
)
