package basic

import (
	"micro_demo/basic/config"
	"github.com/micro/go-plugins/config/source/grpc/v2"
	"github.com/micro/go-micro/v2/registry"
	"micro_demo/basic/common"
	"fmt"
)

var (
	pluginFuncs []func()
)

// Options ...
type Options struct {
	EnableDB    bool
	EnableRedis bool
	cfgOps      []config.Option
}

// Option ...
type Option func(o *Options)

// Init 初始化配置文件
func Init(opts ...config.Option) {
	// 初始化配置
	config.Init(opts...)

	// 加载依赖配置的插件
	for _, f := range pluginFuncs {
		f()
	}
}

// Register 注册方法
func Register(f func()) {
	pluginFuncs = append(pluginFuncs, f)
}



// RegistryOptions  ...
func RegistryOptions(ops *registry.Options) {
	etcdCfg := &common.Etcd{}
	err := config.C().App("etcd", etcdCfg)
	if err != nil {
		panic(err)
	}
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.Host, etcdCfg.Port)}
}



// InitCfg 初始化配置文件 
func InitCfg(appName string, cfg interface{}) {
	source := grpc.NewSource(
		grpc.WithAddress("127.0.0.1:9600"),
		grpc.WithPath("micro"),
	)

	Init(config.WithSource(source))

	err := config.C().App(appName, cfg)
	if err != nil {
		panic(err)
	}

	return
}



