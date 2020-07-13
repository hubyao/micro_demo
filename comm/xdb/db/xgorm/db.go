package xgorm

import (
	"sync"

	"micro_demo/basic/config"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro-in-cn/starter-kit/console/account/conf"
	log "github.com/micro/go-micro/v2/logger"

	"micro_demo/basic"
	"time"

	"github.com/jinzhu/gorm"
)

var (
	dbConf conf.Database
	db     *gorm.DB
	once   sync.Once
)

// init 初始化
func init() {
	basic.Register(initDB)
}


type dbCfg struct {
	Mysql Mysql `json："mysql"`
}

// Mysql mySQL 配置
type Mysql struct {
	URL               string `json:"url"`
	Enable            bool   `json:"enabled"`
	MaxIdleConnection int    `json:"maxIdleConnection"`
	MaxOpenConnection int    `json:"maxOpenConnection"`
	ConnMaxLifetime time.Duration    `json:"connMaxLifetime"`
}
// Init ...
func initDB() {


	log.Info("[initMysql] 初始化Mysql")

	c := config.C()
	cfg := &dbCfg{}

	err := c.App("db", cfg)
	if err != nil {
		log.Warnf("[initMysql] %s", err)
	}


	dbConf = conf.Database{}
	xdb, err := gorm.Open("mysql", cfg.Mysql.URL)
	if err != nil {
		log.Fatal(err)
		return
	}

	xdb.DB().SetMaxOpenConns(cfg.Mysql.MaxOpenConnection)
	xdb.DB().SetMaxIdleConns(cfg.Mysql.MaxIdleConnection)
	xdb.DB().SetConnMaxLifetime(cfg.Mysql.ConnMaxLifetime)
	xdb.SingularTable(true)
	db = xdb
}

// GetDB 获取db
func GetDB() *gorm.DB {
	return db
}
