package xgorm

import (
	"fmt"
	"sync"

	"micro_demo/basic/config"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/jinzhu/gorm"
	"github.com/micro-in-cn/starter-kit/console/account/conf"
	"github.com/micro/go-micro/v2/util/log"
)

var (
	dbConf conf.Database
	db     *gorm.DB
	once   sync.Once
)

// Init ...
func Init() {
	dbConf = conf.Database{}
	log.Debugf("config %v", config.GetMysqlConfig().GetURL())

	fmt.Println("config ", config.GetMysqlConfig().GetURL())
	xdb, err := gorm.Open("mysql", config.GetMysqlConfig().GetURL())
	if err != nil {
		log.Fatal(err)
		return
	}

	xdb.DB().SetMaxOpenConns(dbConf.MaxOpenConns)
	xdb.DB().SetMaxIdleConns(dbConf.MaxIdleConns)
	xdb.DB().SetConnMaxLifetime(dbConf.ConnMaxLifetime)
	xdb.SingularTable(true)

	db = xdb
}

// GetDB 获取db
func GetDB() *gorm.DB {

	return db
}
