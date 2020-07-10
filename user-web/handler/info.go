package handler

import (
	"context"
	"micro_demo/comm/xhttp"
	bpUser "micro_demo/proto/user"

	log "github.com/micro/go-micro/v2/logger"

	"github.com/gin-gonic/gin"
)

// QueryUserByName 查询名字
func QueryUserByName(c *gin.Context) {
	req := QueryUserByNameReq{}
	rsp := QueryUserByNameRsp{}

	// 绑定数据
	if err := c.ShouldBind(&req); err != nil {
		xhttp.SendJsonResponse(c, err, nil)
		return
	}

	// 调用后台服务
	data, err := UserPbClient.QueryUserByName(context.TODO(), &bpUser.Request{
		UserName: req.UserName,
	})

	if err != nil {
		log.Fatal(err)
		xhttp.SendJsonResponse(c, err, nil)
		return
	}

	rsp.Pwd = data.User.Pwd
	xhttp.SendJsonResponse(c, nil, rsp)
}

// QueryUserByNameReq ...
type QueryUserByNameReq struct {
	UserName string `form:"user_name" json:"user_name"`
}

type QueryUserByNameRsp struct {
	Pwd string `json:"pwd"`
}
