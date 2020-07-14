module micro_demo/user-web

go 1.13

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis v6.15.8+incompatible // indirect
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro v1.18.1-0.20200110215436-f50a50eeb36a
	github.com/micro/go-micro/v2 v2.9.1
	micro_demo/basic v0.0.0-00010101000000-000000000000
	micro_demo/comm v0.0.0-00010101000000-000000000000
	micro_demo/proto v0.0.0-00010101000000-000000000000
)

replace (
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
	micro_demo/basic => ../basic
	micro_demo/comm => ../comm
	micro_demo/proto => ../proto
)
