package basic

import (
	"micro_demo/basic/config"
	"micro_demo/basic/db"
	"micro_demo/basic/db/xgorm"
)

func Init() {
	config.Init()
	db.Init()
	xgorm.Init()
	// redis.Init()
}
