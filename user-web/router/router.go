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

	u := g.Group("/v1/user")
	{
		u.POST("sms",handler.Sms) 					 // 发送验证码
		u.POST("/phone/login",handler.PhoneLogin) 	 // 手机号登陆
		u.POST("/wechat/login",handler.WechatLogin)    // 微信登陆
	}
}