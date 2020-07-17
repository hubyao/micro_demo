module micro_demo/user-srv

go 1.13

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	micro_demo/basic v0.0.0-00010101000000-000000000000
	micro_demo/proto v0.0.0-00010101000000-000000000000
)

replace (
	micro_demo/basic => ../basic
	micro_demo/comm => ../comm
	micro_demo/proto => ../proto
)
