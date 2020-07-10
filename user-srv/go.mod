module micro_demo/user-srv

go 1.13

require (
	github.com/go-redis/redis v6.15.8+incompatible // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/jinzhu/gorm v1.9.14 // indirect
	github.com/micro-in-cn/starter-kit v1.18.0 // indirect
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	micro_demo/basic v0.0.0-00010101000000-000000000000
	micro_demo/proto v0.0.0-00010101000000-000000000000
)

replace (
	micro_demo/basic => ../basic
	micro_demo/proto => ../proto
)
