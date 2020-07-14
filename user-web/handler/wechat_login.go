package handler

import (
	"github.com/gin-gonic/gin"
	"micro_demo/comm/xhttp"
)

// 微信登陆
func WechatLogin(c *gin.Context) {



	xhttp.SendJsonResponse(c, err, rsp)
}