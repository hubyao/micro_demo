package handler

import (
	"github.com/gin-gonic/gin"
	"micro_demo/comm/xhttp"
)

// ad 广告
func Ad(c *gin.Context)  {
	rsp:= adRsp{}
	rsp = []adList{
		{
			Name:    "广告一",
			JumpUrl: "https://baidu.com",
			ImgUrl:  "https://goss.veer.com/creative/vcg/veer/800water/veer-302386254.jpg",
		},
		{
			Name:    "广告二",
			JumpUrl: "https://baidu.com",
			ImgUrl:  "https://goss.veer.com/creative/vcg/veer/800water/veer-302386254.jpg",
		},
		{
			Name:    "广告三",
			JumpUrl: "https://baidu.com",
			ImgUrl:  "https://goss.veer.com/creative/vcg/veer/800water/veer-302386254.jpg",
		},
	}
	xhttp.OkRsp(c,rsp)
}

type adRsp []adList

type adList struct {
	Name string `json:"name"`
	JumpUrl  string  `json:"jump_url"` // 跳转链接
	ImgUrl string  `json:"img_url"` // 图片链接
}
