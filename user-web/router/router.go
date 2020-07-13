package router

import (

	"micro_demo/user-web/handler"
	"github.com/gin-gonic/gin"
	"micro_demo/comm/xhttp/middleware"
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
		u.GET("name", handler.GetName)
		u.GET("info", handler.QueryUserByName)
	}
}