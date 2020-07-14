/*
 * @Author       : jianyao
 * @Date         : 2020-07-14 08:32:22
 * @LastEditTime : 2020-07-14 08:38:23
 * @Description  : file content
 */ 
package xgorm

import (
	"sync"

	"micro_demo/basic/config"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/jinzhu/gorm"
	"github.com/micro-in-cn/starter-kit/console/account/conf"
	"github.com/micro/go-micro/v2/util/log"
)

var (
	dbConf conf.Database // 数据库配置
	db     *gorm.DB      // gorm cli
	once   sync.Once     // 用于单例

	ErrRecordNotFound = gorm.ErrRecordNotFound

)

// Init 初始化
func Init() {
	// 单例
	once.Do(func() {
		dbConf = conf.Database{}
		xdb, err := gorm.Open("mysql", config.GetMysqlConfig().GetURL())
		if err != nil {
			log.Fatal(err)
			return
		}

		// 设置连接池
		xdb.DB().SetMaxOpenConns(dbConf.MaxOpenConns)
		xdb.DB().SetMaxIdleConns(dbConf.MaxIdleConns)
		xdb.DB().SetConnMaxLifetime(dbConf.ConnMaxLifetime)
		xdb.SingularTable(true)
		db = xdb
	})
}

// GetDB 获取db
func GetDB() *gorm.DB {
	return db
}
