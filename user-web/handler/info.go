package handler

import (
	"context"
	"fmt"
	"log"
	"micro_demo/comm/xhttp"
	bpUser "micro_demo/proto/user"

	"github.com/gin-gonic/gin"
)

// QueryUserByNameReq ...
type QueryUserByNameReq struct {
	UserName string `form:"user_name" json:"user_name"`
}

type QueryUserByNameRsp struct {
	Pwd string `json:"pwd"`
}

// QueryUserByName 查询名字
func QueryUserByName(c *gin.Context) {
	req := QueryUserByNameReq{}
	rsp := QueryUserByNameRsp{}

	// 绑定数据
	if err := c.ShouldBind(&req); err != nil {
		xhttp.SendJsonResponse(c, err, nil)
		return
	}

	fmt.Println("req ", req)

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
