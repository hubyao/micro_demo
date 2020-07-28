module micro_demo/basic

go 1.13

require (
	github.com/go-redis/redis v6.15.8+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/jinzhu/gorm v1.9.14
	github.com/micro-in-cn/starter-kit v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	micro_demo/comm v0.0.0-00010101000000-000000000000
)

replace (
	micro_demo/basic => ../basic
	micro_demo/comm => ../comm
	micro_demo/proto => ../proto
)
