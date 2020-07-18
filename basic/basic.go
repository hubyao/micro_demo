package basic

import (
	"micro_demo/basic/config"
	"micro_demo/basic/db"
	"micro_demo/basic/db/xgorm"
	"micro_demo/basic/redis"
)

func Init() {
	config.Init()
	db.Init()
	xgorm.Init()
	redis.Init()
}
