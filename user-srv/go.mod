module micro_demo/user-srv

go 1.13

require (
	github.com/go-redis/redis v6.15.8+incompatible // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gookit/color v1.2.6 // indirect
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/xxjwxc/public v0.0.0-20200710160137-ccf3e4f07a03 // indirect
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/tools v0.0.0-20200713011307-fd294ab11aed // indirect
	honnef.co/go/tools v0.0.1-2020.1.4 // indirect
	micro_demo/basic v0.0.0-00010101000000-000000000000
	micro_demo/comm v0.0.0-00010101000000-000000000000
	micro_demo/proto v0.0.0-00010101000000-000000000000
)

replace (
	micro_demo/basic => ../basic
	micro_demo/comm => ../comm
	micro_demo/proto => ../proto
)
