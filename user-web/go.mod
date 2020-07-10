module micro_demo/user-web

go 1.13

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis v6.15.8+incompatible // indirect
	github.com/jinzhu/gorm v1.9.14 // indirect
	github.com/micro-in-cn/starter-kit v1.18.0 // indirect
	github.com/micro/go-micro/v2 v2.9.1
	github.com/natefinch/lumberjack v2.0.0+incompatible // indirect
	micro_demo/basic v0.0.0-00010101000000-000000000000 // indirect
	micro_demo/comm v0.0.0-00010101000000-000000000000
	micro_demo/proto v0.0.0-00010101000000-000000000000
)

replace (
	micro_demo/basic => ../basic
	micro_demo/comm => ../comm
	micro_demo/proto => ../proto
)
