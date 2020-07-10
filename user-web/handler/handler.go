package handler

import (
	"fmt"
	"micro_demo/comm/xhttp"
	bpUser "micro_demo/proto/user"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
)

// Error 错误结构体
type Error struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

// GetNameReq 获取名字
type GetNameReq struct {
	Name string `form:"name" json:"name"`
	Age  int    `form:"age" json:"age"`
}

// GetNameRsp 响应
type GetNameRsp struct {
	Name string `json:"cname"`
}

var (
	UserPbClient bpUser.UserService
)

func Init() {
	UserPbClient = bpUser.NewUserService("mu.micro.book.srv.user", client.DefaultClient)
}

// GetName ...
func GetName(c *gin.Context) {
	req := GetNameReq{}
	rsp := GetNameRsp{}

	// 绑定数据
	if err := c.ShouldBind(&req); err != nil {
		xhttp.SendJsonResponse(c, err, nil)
		// return
	}

	fmt.Println("req ", req)
	rsp.Name = req.Name

	xhttp.SendJsonResponse(c, nil, rsp)
}
