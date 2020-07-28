package xgorm

import (
	"log"
	//"micro_demo/comm/logging"
	"sync"

	"micro_demo/basic/config"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/jinzhu/gorm"
	"github.com/micro-in-cn/starter-kit/console/account/conf"
	//"github.com/micro/go-micro/v2/util/log"
)

var (
	dbConf conf.Database // 数据库配置
	gdb    *gorm.DB      // gorm cli
	once   sync.Once     // 用于单例
)

// Init 初始化
func Init() {
	// 单例
	once.Do(func() {

		mysqlCfg := &config.MysqlConfig{}
		if !mysqlCfg.Enable {
			return
		}

		config.GetConfig(mysqlCfg)

		dbConf = conf.Database{
			MaxOpenConns:    mysqlCfg.MaxOpenConnection,
			MaxIdleConns:    mysqlCfg.MaxIdleConnection,
			ConnMaxLifetime: mysqlCfg.ConnMaxLifetime,
		}

		cli := mysqlCfg.GetURL()
		xdb, err := gorm.Open("mysql", cli)
		if err != nil {
			log.Fatal(err)
			return
		}

		// 设置连接池
		xdb.DB().SetMaxOpenConns(dbConf.MaxOpenConns)
		xdb.DB().SetMaxIdleConns(dbConf.MaxIdleConns)
		xdb.DB().SetConnMaxLifetime(dbConf.ConnMaxLifetime)
		xdb.SingularTable(true)
		gdb = xdb
	})
}

// GetDB 获取db
func GetDB() *gorm.DB {
	return gdb
}
