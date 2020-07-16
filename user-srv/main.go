/*
 * @Author       : jianyao
 * @Date         : 2020-07-16 07:24:59
 * @LastEditTime : 2020-07-16 08:13:35
 * @Description  : file content
 */ 
package main

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"micro_demo/comm/micro/allenxuxu/tracer"

	"github.com/micro/cli/v2"
	micro "github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"micro_demo/basic"
	"micro_demo/basic/config"
	s "micro_demo/proto/user"
	"micro_demo/user-srv/handler"
	"micro_demo/user-srv/model"
	openTrace "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"

)

func main() {
	// 初始化配置、数据库等信息
	basic.Init()

	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)


	t, io, err := tracer.NewTracer("mu.micro.book.srv.user", "localhost:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()

	opentracing.SetGlobalTracer(t)

	// New Service
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.user"),
		micro.Registry(micReg),
		micro.Version("latest"),
		micro.WrapHandler(openTrace.NewHandlerWrapper(opentracing.GlobalTracer())),
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

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
