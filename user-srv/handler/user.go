package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"
	pbUser "micro_demo/proto/user"
	modelUser "micro_demo/user-srv/model/user"
)

type Service struct{}

var (
	userService us.Service
)

// Init 初始化handler
func Init() {
	var err error
	userService, err = us.GetService()
	if err != nil {
		log.Fatal("[Init] 初始化Handler错误")
		return
	}
}

// QueryUserByName 通过参数中的名字返回用户
func (e *Service) QueryUserByName(ctx context.Context, req *pbUser.AddUserReq, rsp *pbUser.AddUserRsp) error {
	err := userService.AddUser(modelUser.User{
		
	})
	

	// user, err := userService.QueryUserByName(req.UserName)
	// if err != nil {
	// 	rsp.BaseResponse.Success = false
	// 	rsp.BaseResponse.Error= &s.Error{
	// 		Code:   500,
	// 		Message: err.Error(),
	// 	}

	// 	return nil
	// }

	// rsp.User = user
	// rsp.BaseResponse.Success = true
	// return nil
}
