package handler

import (
	"github.com/gin-gonic/gin"
	"micro_demo/comm/xhttp"
)

// IncentiveVideo  激励视频列表
func IncentiveVideo(c *gin.Context){
	//req := incentiveVideoReq{}
	rsp := incentiveVideoRep{}

	rsp = []incentiveVideo{
		{
			TotalNum: 5,
			PlayNum:  0,
		},
		{
			TotalNum: 5,
			PlayNum:  0,
		},
		{
			TotalNum: 5,
			PlayNum:  0,
		},
		{
			TotalNum: 5,
			PlayNum:  0,
		},
		{
			TotalNum: 5,
			PlayNum:  0,
		},
	}

	xhttp.OkRsp(c,rsp)
}

type incentiveVideoReq struct {
}

type incentiveVideoRep []incentiveVideo

type incentiveVideo struct {
	TotalNum int `json:"total_num"` // 总播放次数
	PlayNum  int `json:"play_num"`  // 已播放次数
}
