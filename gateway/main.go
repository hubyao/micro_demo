package main

import (
	cors "github.com/micro/go-plugins/micro/cors/v2"
	"github.com/micro/micro/v2/cmd"
	"github.com/micro/micro/v2/plugin"
)

const name = "API gateway"

// 在此注册使用的插件
func main() {
	// 注册插件
	plugin.Register(cors.NewPlugin())


	//plugin.Register(plugin.NewPlugin(
	//	plugin.WithHandler(xlog()),
	//))


	cmd.Init()
}


// // 日志处理
// func xlog() plugin.Handler {
// 	return func(h http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		})
// 	}
// }
