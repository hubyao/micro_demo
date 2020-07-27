/*
 * @Author       : jianyao
 * @Date         : 2020-07-21 06:22:54
 * @LastEditTime : 2020-07-21 06:53:49
 * @Description  : file content
 */
package router

import (
	"micro_demo/comm/xhttp/middleware"
	"micro_demo/user-web/handler"

	"github.com/gin-gonic/gin"
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
	sl := g.Group("/sportlottery/v1")
	{
		sl.GET("/jackpot", handler.ToDayJackpot) // 今日奖池
		sl.GET("/rule", handler.Rule)            // 规则
		sl.GET("ad", handler.Ad)                 // 广告
	}

}
