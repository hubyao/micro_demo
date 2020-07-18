package handler

import (
	"github.com/gin-gonic/gin"
	"micro_demo/comm/xhttp"
)

// Rule 玩法规则
func Rule(c *gin.Context)  {
	rsp :=ruleRsp{}
	rsp.Data = "玩法规则"
	xhttp.OkRsp(c,rsp)
}

type  ruleRsp struct {
	Data string `json:"data"`
}