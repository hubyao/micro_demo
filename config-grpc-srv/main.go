package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
	"flag"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/file"
	log "github.com/micro/go-micro/v2/logger"
	proto "github.com/micro/go-plugins/config/source/grpc/v2/proto"
	grpc2 "google.golang.org/grpc"
	"micro_demo/basic/common"
	"os"
	// "micro_demo/basic/common"
)

var (
	mux        sync.RWMutex
	configMaps = make(map[string]*proto.ChangeSet)
	// 根据需要改成可配置的app列表
	apps = []string{"micro_loc"}
)

var (
	env string  // 环境
)

// Service ...
type Service struct{}


// flagParse 参数解析
func flagParseEnv(){
	// 获取命令行参数
    flag.StringVar(&env, "env", "", "环境, 可选项 loc,dev,prod")
	flag.Parse()
	
	if env == ""{
		env = common.EnvLoc
	}
	
	switch env {
		case common.EnvLoc:
			apps = []string{"micro_loc"}
		case common.EnvDev:
			apps = []string{"micro_dev"}
		case common.EnvProd:
			apps = []string{"micro_prod"}
		default:
			fmt.Println("env只支持 loc,dev,prod")
			os.Exit(1)
	}
	
	log.Infof("启动环境: %v",env)

	
}


func main() {
	// 解析环境
	flagParseEnv()

	// 灾难恢复
	defer func() {
		if r := recover(); r != nil {
			log.Infof("[main] Recovered in f %v", r)
		}
	}()

	// 加载并侦听配置文件
	err := loadAndWatchConfigFile()
	if err != nil {
		log.Fatal(err)
	}

	// 新建grpc Server服务
	service := grpc2.NewServer()
	proto.RegisterSourceServer(service, new(Service))
	ts, err := net.Listen("tcp", ":9600")
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("configServer started")

	// 启动
	err = service.Serve(ts)
	if err != nil {
		log.Fatal(err)
	}
}

// Read 读取文件
func (s Service) Read(ctx context.Context, req *proto.ReadRequest) (rsp *proto.ReadResponse, err error) {
	appName := parsePath(req.Path)

	rsp = &proto.ReadResponse{
		ChangeSet: getConfig(appName),
	}
	return
}

// Watch 监听文件
func (s Service) Watch(req *proto.WatchRequest, server proto.Source_WatchServer) (err error) {
	appName := parsePath(req.Path)
	rsp := &proto.WatchResponse{
		ChangeSet: getConfig(appName),
	}
	if err = server.Send(rsp); err != nil {
		log.Errorf("[Watch] 侦听处理异常，%s", err)
		return err
	}

	return
}

// loadAndWatchConfigFile 加载文件
func loadAndWatchConfigFile() (err error) {
	// 加载每个应用的配置文件
	for _, app := range apps {
		if err := config.Load(file.NewSource(
			file.WithPath("./conf/" + app + ".yml"),
		)); err != nil {
			log.Fatalf("[loadAndWatchConfigFile] 加载应用配置文件 异常，%s", err)
			return err
		}
	}

	// 侦听文件变动
	watcher, err := config.Watch()
	if err != nil {
		log.Fatalf("[loadAndWatchConfigFile] 开始侦听应用配置文件变动 异常，%s", err)
		return err
	}

	go func() {
		for {
			v, err := watcher.Next()
			if err != nil {
				log.Fatalf("[loadAndWatchConfigFile] 侦听应用配置文件变动 异常， %s", err)
				return
			}

			log.Infof("[loadAndWatchConfigFile] 文件变动，%s", string(v.Bytes()))
		}
	}()

	return
}

// getConfig 获取文件
func getConfig(appName string) *proto.ChangeSet {
	bytes := config.Get(appName).Bytes()

	log.Infof("[getConfig] appName：%s", appName)
	return &proto.ChangeSet{
		Data:      bytes,
		Checksum:  fmt.Sprintf("%x", md5.Sum(bytes)),
		Format:    "yml",
		Source:    "file",
		Timestamp: time.Now().Unix()}
}


// parsePath 解析地址
func parsePath(path string) (appName string) {
	paths := strings.Split(path, "/")

	if paths[0] == "" && len(paths) > 1 {
		return paths[1]
	}

	return paths[0]
}
