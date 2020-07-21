/*
 * @Author       : jianyao
 * @Date         : 2020-07-21 06:22:54
 * @LastEditTime : 2020-07-21 06:53:56
 * @Description  : file content
 */
package main

import (
	"fmt"

	"micro_demo/basic"
	"micro_demo/basic/common"
	"micro_demo/basic/config"
	"micro_demo/user-web/handler"
	"micro_demo/user-web/router"

	"github.com/gin-gonic/gin"
	"github.com/micro/cli/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
)

var (
	appName    = "sportlottery"
	appAllName = common.AppNamePrefix + ".api." + appName
)

func main() {

	// 初始化配置
	basic.Init()

	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)

	// 创建新服务
	service := web.NewService(
		// 后面两个web，第一个是指是web类型的服务，第二个是服务自身的名字
		web.Name(appAllName),
		web.Version("latest"),
		web.Registry(micReg),
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

	g := gin.Default()

	router.Load(
		g,
	)

	service.Handle("/", g)

	if err := service.Run(); err != nil {
		log.Fatal("", err)
	}

	// 运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
