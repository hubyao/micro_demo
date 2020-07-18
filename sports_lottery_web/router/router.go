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

	// 体彩
	sl := g.Group("/v1/sports_lottery")
	{
		sl.GET("/jackpot", handler.ToDayJackpot) // 今日奖池
		sl.GET("/rule", handler.Rule)            // 规则
		sl.GET("ad", handler.Ad)                 // 广告
	}

}
