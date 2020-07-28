package main

import (
	"github.com/micro/cli/v2"
	micro "github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry/etcd"
	"micro_demo/basic"
	"micro_demo/basic/common"
	"micro_demo/basic/config"
	s "micro_demo/proto/user"
	"micro_demo/user-srv/handler"
	"micro_demo/user-srv/model"
)

var (
	appFullName = ""
)

func main() {
	// 初始化配置、数据库等信息
	basic.Init()
	appFullName = common.AppNamePrefix + ".srv." + config.GetProfiles().AppName
	// 使用etcd注册
	micReg := etcd.NewRegistry(config.RegistryOptions)

	// New Service
	service := micro.NewService(
		micro.Name(appFullName),
		micro.Registry(micReg),
		micro.Version("latest"),
		micro.Flags(
			&cli.StringFlag{
				Name:  "env",
				Usage: "",
				Value: "",
			}),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) error {
			// 初始化模型层
			model.Init()
			// 初始化handler
			handler.Init()

			return nil
		}),
	)

	// 注册服务
	s.RegisterUserHandler(service.Server(), new(handler.Service))

	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
