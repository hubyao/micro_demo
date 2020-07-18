package handler

import (
	"github.com/gin-gonic/gin"
	"micro_demo/comm/xhttp"
)

// ToDayJackpot 今日奖池
func ToDayJackpot(c *gin.Context)  {
	rsp := toDayJackpotRsp{}

	rsp.Num = 13292033
	rsp.CurrentLotteryDate = 1594991944
	rsp.LastLotteryNum ="12.07.09.13"
	xhttp.OkRsp(c,rsp)
}

type toDayJackpotReq struct {

}

type toDayJackpotRsp struct {
	Num int64 `json:"num"` // 今日奖池累计数
	CurrentLotteryDate int64 `json:"current_lottery_date"`// 本期开奖时间秒
	LastLotteryNum string `json:"last_lottery_num"` // 上期毫秒

}
