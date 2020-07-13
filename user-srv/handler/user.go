package handler

import (
	"context"
	"micro_demo/comm/logging"
	log "github.com/micro/go-micro/v2/logger"
	pbUser "micro_demo/proto/user"
	modelUser "micro_demo/user-srv/model/user"
)

type Service struct{}


var (
	userService modelUser.Service
)

// Init 初始化handler
func Init() {
	var err error
	userService, err = modelUser.GetService()
	if err != nil {
		log.Fatal("[Init] 初始化Handler错误")
		return
	}
}

// AddUser 添加用户信息
func (e *Service) AddUser(ctx context.Context, req *pbUser.AddUserReq, rsp *pbUser.AddUserRsp) error {
	rsp = &pbUser.AddUserRsp{
		BaseResponse: &pbUser.BaseResponse{
			Success: true,
			Error:   nil,
		},
	}

	err := userService.AddUser(modelUser.User{
		Phone:      req.Phone,
		Nick:       req.Nick,
	})

	if err != nil {
		logging.Logger().Error(err)
		rsp.BaseResponse.Success = false
		rsp.BaseResponse.Error = &pbUser.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	return nil
}


// GetFromUid 根据用户id获取信息
func (e *Service) GetFromUid(ctx context.Context, req *pbUser.GetFromUidReq, rsp *pbUser.GetFromUidRsp) error {
	rsp = &pbUser.GetFromUidRsp{
		BaseResponse: &pbUser.BaseResponse{
			Success: true,
			Error:   nil,
		},
		UserInfo: nil,
	}

	data,err := userService.GetFromUId(req.Uid)
	if err != nil {
		logging.Logger().Error(err)
		rsp.BaseResponse.Success = false
		rsp.BaseResponse.Error = &pbUser.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	if nil == data{
		return  nil
	}


	rsp.UserInfo = &pbUser.UserInfo{
		Uid:   data.UId,
		Nick:  data.Nick,
		Phone: data.Phone,
	}

	return nil
}

// GetFromPhone 根据手机号获取信息
func (e *Service) GetFromPhone(ctx context.Context, req *pbUser.GetFromPhoneReq, rsp *pbUser.GetFromPhoneRsp) error {
	rsp = &pbUser.GetFromPhoneRsp{
		BaseResponse: &pbUser.BaseResponse{
			Success: true,
			Error:   nil,
		},
		UserInfo: nil,
	}

	data,err := userService.GetFromPhone(req.Phone)
	if err != nil {
		logging.Logger().Error(err)
		rsp.BaseResponse.Success = false
		rsp.BaseResponse.Error = &pbUser.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	if nil == data{
		return  nil
	}


	rsp.UserInfo = &pbUser.UserInfo{
		Uid:   data.UId,
		Nick:  data.Nick,
		Phone: data.Phone,
	}

	return nil
}

// UpdateUser 更新用户信息
func (e *Service) UpdateUser(ctx context.Context, req *pbUser.UpdateUserReq, rsp *pbUser.UpdateUserRsp) error {
	rsp = &pbUser.UpdateUserRsp{
		BaseResponse: &pbUser.BaseResponse{
			Success: true,
			Error:   nil,
		},
	}

	err := userService.AddUser(modelUser.User{
		Nick:       req.Nick,
	})

	if err != nil {
		logging.Logger().Error(err)
		rsp.BaseResponse.Success = false
		rsp.BaseResponse.Error = &pbUser.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	return nil
}