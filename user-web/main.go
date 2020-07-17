package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/cli/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	opentracing "github.com/opentracing/opentracing-go"
	"micro_demo/basic"
	"micro_demo/basic/config"
	"micro_demo/comm/micro/allenxuxu/tracer"
	"micro_demo/comm/micro/allenxuxu/wrapper/tracer/opentracing/gin2micro"
	"micro_demo/user-web/handler"
)

func main() {

	// 初始化配置
	basic.Init()

	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)

	gin2micro.SetSamplingFrequency(50)


	t, io, err := tracer.NewTracer("mu.micro.book.web.user", "")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)


	// 创建新服务
	service := web.NewService(
		// 后面两个web，第一个是指是web类型的服务，第二个是服务自身的名字
		web.Name("mu.micro.book.web.user"),
		web.Version("latest"),
		web.Registry(micReg),
		web.Address(":8088"),
		web.Action(
			func(c *cli.Context) {
				// 初始化handler
				handler.Init()
			}),
	)

	// 初始化服务
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	gin.SetMode("debug")

	g := gin.Default()

	r := g.Group("/")
	r.Use(gin2micro.TracerWrapper)

	//service.Handle("/", g)
	// 开启链路追踪插件
	service.Handle("/", g)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal("", err)
	}

}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
