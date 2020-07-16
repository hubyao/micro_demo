package router

import (
	"github.com/gin-gonic/gin"
	"micro_demo/comm/micro/tracer/gin2micro"
	"micro_demo/comm/xhttp/middleware"
	"micro_demo/user-web/handler"
)

// Load 加载中间件
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(mw...)
	g.Use(middleware.DetailLogger())
	g.Use(gin2micro.TracerWrapper)
	Handle(g)

	return g
}

// Handle ...
func Handle(g *gin.Engine) {

	u := g.Group("/v1/user")
	{
		u.GET("name", handler.GetName)
		u.GET("info", handler.QueryUserByName)
	}
}
