package handler

import (
	bpUser "micro_demo/proto/user"

	"github.com/micro/go-micro/v2/client"
)


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


