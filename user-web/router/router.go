/*
 * @Author       : jianyao
 * @Date         : 2020-07-14 08:32:22
 * @LastEditTime : 2020-07-14 08:36:54
 * @Description  : file content
 */
package router

import (
	"github.com/gin-gonic/gin"
	"micro_demo/comm/xhttp/middleware"
	"micro_demo/user-web/handler"
)

// Load 加载中间件
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(mw...)
	g.Use(middleware.DetailLogger())

	Handle(g)

	return g
}

// Handle ...
func Handle(g *gin.Engine) {
	// 需要登录验证的
	_Authed := g.Group("", middleware.Auth())

	// 不需要登录
	e := g.Group("/v1/user")
	{
		e.POST("/sms", handler.Sms) // 发送验证码

		e.POST("/phone/login", handler.PhoneLogin)   // 手机号登陆
		e.POST("/wechat/login", handler.WechatLogin) //  微信登陆

		e.POST("/register", handler.UserRegister) // 注册
	}


	// 需要登录的接口
	u := _Authed.Group("/v1/user")
	{
		// 好友助力
		u.GET("/friend/help", handler.FriendHelp) // 获取好友助力列表

		// 每日任务
		u.GET("/daily_task", handler.DailyTask) // 获取每日任务列表

		// 激励视频
		u.GET("/incentive_video", handler.IncentiveVideo) // 获取激励视频列表
	}

	// 广告

	//
}
