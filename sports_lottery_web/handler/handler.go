package handler

import (
	bpUser "micro_demo/proto/user"

	"github.com/micro/go-micro/v2/client"
)



var (
	UserPbClient bpUser.UserService
)

func Init() {
	UserPbClient = bpUser.NewUserService("mu.micro.book.srv.user", client.DefaultClient)
}

