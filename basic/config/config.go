package config

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/micro/go-micro/v2/config"
	//"github.com/micro/go-micro/v2/config/source"
	"github.com/micro/go-micro/v2/config/source/file"
	log "github.com/micro/go-micro/v2/logger"
)

var (
	err error
)

var (
	defaultRootPath         = "app"
	defaultConfigFilePrefix = "application-"
	defaultConfigFile       = "application-loc"
	profiles                defaultProfiles
	m                       sync.RWMutex
	inited                  bool
	configSrv               *configService
	env                     string
)

type configService struct {
	initField map[string]IGetConfig
}

// TODO: 解析与micro原生命令行参数出现冲突无法实现
// flagParse 参数解析
func flagParseEnv() {
	//// 获取命令行参数
	//flag.StringVar(&env, "env", "", "环境, 可选项 loc,dev,prod")
	////flag.Parse()
	////flag.
	//
	//if env == "" {
	//	env = common.EnvLoc
	//}
	//
	//switch env {
	//case common.EnvLoc:
	//	defaultConfigFile = "application-loc"
	//case common.EnvDev:
	//	defaultConfigFile = "application-dev"
	//case common.EnvProd:
	//	defaultConfigFile = "application-prod"
	//default:
	//	fmt.Println("env只支持 loc,dev,prod")
	//	os.Exit(1)
	//}
	//
	//log.Infof("启动环境: %v", env)

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

	// 找到需要引入的新配置文件
	if err = config.Get(defaultRootPath, "profiles").Scan(&profiles); err != nil {
		panic(err)
	}

	log.Infof("[Init] 加载配置文件：path: %s, %+v\n", filePath, profiles)

	// 开始导入新文件
	//if len(profiles.GetInclude()) > 0 {

	sources := file.NewSource(file.WithPath(filePath))

	// 加载include的文件
	if err = config.Load(sources); err != nil {
		panic(err)
	}
	//}

	// 赋值
	//config.Get(defaultRootPath, "etcd").Scan(&etcdConfig)
	//config.Get(defaultRootPath, "mysql").Scan(&mysqlConfig)
	//config.Get(defaultRootPath, "redis").Scan(&redisConfig)
	//config.Get(defaultRootPath, "jwt").Scan(&jwtConfig)

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
