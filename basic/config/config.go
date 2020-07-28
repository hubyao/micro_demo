package config

import (
	"flag"
	"fmt"
	"micro_demo/basic/common"
	"micro_demo/comm/logging"
	"os"
	"path/filepath"
	"sync"

	"github.com/micro/go-micro/v2/config"
	//"github.com/micro/go-micro/v2/config/source"
	"github.com/micro/go-micro/v2/config/source/file"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
)

var (
	err error
)

var (
	defaultRootPath   = "app"
	defaultConfigFile = "application-loc"
	m                 sync.RWMutex
	inited            bool
	configSrv         *configService
	env               string
)

type configService struct {
	initField map[string]IGetConfig
}

// TODO: 解析与micro原生命令行参数出现冲突无法实现
// flagParse 参数解析
func flagParseEnv() {
	// 获取命令行参数
	flag.StringVar(&env, "env", "", "环境, 可选项 loc,dev,prod")
	flag.Parse()

	if env == "" {
		env = common.EnvLoc
	}

	switch env {
	case common.EnvLoc:
		defaultConfigFile = "application-loc"
	case common.EnvDev:
		defaultConfigFile = "application-dev"
	case common.EnvProd:
		defaultConfigFile = "application-prod"
	default:
		fmt.Println("env只支持 loc,dev,prod")
		os.Exit(1)
	}

	log.Infof("启动环境: %v", env)

}

// Init 初始化配置
func Init() {
	m.Lock()
	defer m.Unlock()

	// 解析命令行
	flagParseEnv()

	if inited {
		log.Info("[Init] 配置已经初始化过")
		return
	}

	// 加载yml配置
	// 先加载基础配置
	appPath, _ := filepath.Abs(filepath.Dir(filepath.Join("./", string(filepath.Separator))))

	pt := filepath.Join(appPath, "conf")
	os.Chdir(appPath)
	filePath := pt + string(filepath.Separator) + defaultConfigFile + ".yml"

	// 找到application.yml文件
	if err = config.Load(file.NewSource(file.WithPath(filePath))); err != nil {
		panic(err)
	}

	var data interface{}
	config.Get(defaultRootPath).Scan(&data)
	logging.Logger().Infof("配置数据 %v", data)
	// 标记已经初始化
	inited = true
	configSrv = &configService{
		initField: make(map[string]IGetConfig),
	}
}

type IGetConfig interface {
	GetName() string
	//GetData(interface{}) interface{}
}

// 获取配置文件
// IGetConfig 指针
func GetConfig(data IGetConfig) {
	// TODO: 避免重复赋值
	//t, ok := configSrv.initField[data.GetName()] /*如果确定是真实的,则存在,否则不存在 */
	//if ok {
	//	data = data.GetData(t).(IGetConfig)
	//	fmt.Println("已经初始化了")
	//}
	//config.Config().Set()
	config.Get(defaultRootPath, (data).GetName()).Scan(data)
	//configSrv.initField[data.GetName()] = data

	return
}

// GetProfiles 获取应用属性
func GetProfiles() *ProfilesConfig {
	cfg := &ProfilesConfig{}
	GetConfig(cfg)
	return cfg
}

// registryOptions 注册
func RegistryOptions(ops *registry.Options) {
	etcdCfg := &EtcdConfig{}
	GetConfig(etcdCfg)
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
