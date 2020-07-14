package handler

import (
	"context"
	"micro_demo/comm/logging"
	log "github.com/micro/go-micro/v2/logger"
	"micro_demo/comm/xhttp/errno"
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
		Uid: 0,
	}
	data := &modelUser.User{
		Phone:      req.Phone,
		Nick:       req.Nick,
	}

	err := userService.AddUser(data)

	if err != nil {
		logging.Logger().Error(err)
		rsp.BaseResponse.Success = false
		rsp.BaseResponse.Error = &pbUser.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	rsp.Uid = data.UId
	
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

	err := userService.UpdateUser(modelUser.User{
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
	// rsp = &pbUser.GetFromPhoneRsp{
	// 	BaseResponse: &pbUser.BaseResponse{
	// 		Success: true,
	// 		Error:   nil,
	// 	},
	// 	UserInfo: &pbUser.UserInfo{},
	// }
	
	// logging.Logger().Debug(rsp)
	// data,err := userService.GetFromPhone(req.Phone)

	// logging.Logger().Debug(data)
	// if err != nil {
	// 	logging.Logger().Error(err)
	// 	rsp.BaseResponse.Success = false
	// 	rsp.BaseResponse.Error = &pbUser.Error{
	// 		Code:    500,
	// 		Message: err.Error(),
	// 	}
	// }

	// logging.Logger().Debug(rsp)
	// if nil == data{
	// 	logging.Logger().Debug(rsp)
	// 	return  nil
	// }
	// logging.Logger().Debug(rsp)


	
	// rsp.UserInfo = &pbUser.UserInfo{
	// 	Uid:   data.UId,
	// 	Nick:  data.Nick,
	// 	Phone: data.Phone,
	// }
	

	// logging.Logger().Debug(rsp)
	// return nil
}

// ParseToken 生成token 
func (e *Service) ParseToken(ctx context.Context, req *pbUser.ParseTokenReq, rsp *pbUser.ParseTokenRsp) error {
	rsp = &pbUser.ParseTokenRsp{
		BaseResponse: &pbUser.BaseResponse{
			Success: true,
			Error:   nil,
		},
		Uid:0,
	}

	dataJwt ,err := userService.ParseToken(req.Token)
	if err != nil {
			logging.Logger().Error(err)
			rsp.BaseResponse.Success = false
			rsp.BaseResponse.Error = &pbUser.Error{
				Code:    500,
				Message: err.Error(),
			}
	}
	
	rsp.Uid = dataJwt.UserId
	return nil
}



// GenerateToken 生成token 
func (e *Service) GenerateToken(ctx context.Context, req *pbUser.GenerateTokenReq, rsp *pbUser.GenerateTokenRsp) error {
	rsp = &pbUser.GenerateTokenRsp{
		BaseResponse: &pbUser.BaseResponse{
			Success: true,
			Error:   nil,
		},
		Token:"",
	}

	token ,err := userService.GenerateToken(req.Uid,0)
	if err != nil {
			logging.Logger().Error(err)
			rsp.BaseResponse.Success = false
			rsp.BaseResponse.Error = &pbUser.Error{
				Code:    500,
				Message: err.Error(),
			}
	}
	
	rsp.Token = token
	
	return nil
}



// UserOauthLogin 用户授权登陆
func (e *Service) UserOauthLogin(ctx context.Context, req *pbUser.UserOauthLoginReq, rsp *pbUser.UserOauthLoginRsp) error {
	
	rsp = &pbUser.UserOauthLoginRsp{
		BaseResponse: &pbUser.BaseResponse{
			Success: true,
			Error:   nil,
		},
		Uid: 0,
		Token:"",
	}
	
	dataUserOauth ,err := userService.GetUserOauthByPlatformWechat(req.UserOauth.OpenId)
	if err != nil {
		logging.Logger().Error(err)
		rsp.BaseResponse.Success = false
		rsp.BaseResponse.Error = &pbUser.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	// 存在授权
	if nil != dataUserOauth{
		// 生成token
		token ,err := userService.GenerateToken(dataUserOauth.UId,0)
		if err != nil {
			logging.Logger().Error(err)
			rsp.BaseResponse.Success = false
			rsp.BaseResponse.Error = &pbUser.Error{
				Code:    500,
				Message: err.Error(),
			}
		}

		rsp.Token = token
		rsp.Uid = dataUserOauth.UId
		return nil
	}
	

	// 不存在授权
	// 添加用户信息
	dataUser := &modelUser.User{
		Nick:       req.UserOauth.Name,
	}

	err = userService.AddUser(dataUser)
	if err != nil {
		logging.Logger().Error(err)
		rsp.BaseResponse.Success = false
		rsp.BaseResponse.Error = &pbUser.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	// 添加授权信息
	err = userService.AddUserOauthr(&modelUser.UserOauth{
		Platform:    "wechat",
		OpenId:      req.UserOauth.OpenId,
		Unionid:     req.UserOauth.Unionid,
		Sex:         int(req.UserOauth.Sex),
		Name:        req.UserOauth.Name,
		Avatar:      req.UserOauth.Avatar,
		Sessionkey:  req.UserOauth.Sessionkey,
		UId:         dataUser.UId,
	})

	if err != nil {
		logging.Logger().Error(err)
		rsp.BaseResponse.Success = false
		rsp.BaseResponse.Error = &pbUser.Error{
			Code:    int32(errno.ErrUserAddUserOauth.Code),
			Message: err.Error(),
		}
	}


	// 生成token
	token ,err := userService.GenerateToken(dataUserOauth.UId,0)
	if err != nil {
		logging.Logger().Error(err)
		rsp.BaseResponse.Success = false
		rsp.BaseResponse.Error = &pbUser.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	rsp.Token = token
	rsp.Uid = dataUserOauth.UId
	

	return nil
}


