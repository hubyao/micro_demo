package main

import (
	// "fmt"
	"net/http"

	"micro_demo/basic"
	// "micro_demo/basic/config"
	"micro_demo/user-web/handler"
	"micro_demo/user-web/router"

	"github.com/gin-gonic/gin"
	"github.com/micro/cli/v2"
	log "github.com/micro/go-micro/v2/logger"
	// "github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	"micro_demo/basic/common"
	// "github.com/micro/go-plugins/config/source/grpc/v2"
)



var (
	appName = "user_web"
	cfg     = &userCfg{}
)

type userCfg struct {
	common.AppCfg
}

func main() {

	// 初始化配置
	// initCfg()
	basic.InitCfg(appName,cfg)

	// 使用etcd注册
	micReg := etcd.NewRegistry(basic.RegistryOptions)

	// 创建新服务
	service := web.NewService(
		// 后面两个web，第一个是指是web类型的服务，第二个是服务自身的名字
		web.Name(cfg.Name),
		web.Version(cfg.Version),
		web.Registry(micReg),
		web.Address(cfg.Addr()),
	)

	// 初始化服务
	if err := service.Init(
		web.Action(
			func(c *cli.Context) {
				// 初始化handler
				handler.Init()
			}),
	); err != nil {
		log.Fatal(err)
	}

	gin.SetMode("debug")

	g := gin.New()
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
		return
	})

	router.Load(
		g,
	)

	service.Handle("/", g)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal("", err)
	}

	// 运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}