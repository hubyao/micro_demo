package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"micro_demo/comm/logging"
	"micro_demo/comm/xhttp"
	"micro_demo/comm/xhttp/errno"
	pbUser "micro_demo/proto/user"
)

// PhoneLogin 手机号登陆
func PhoneLogin(c *gin.Context) {
	req := PhoneLoginReq{}
	rsp := PhoneLoginRsp{}
	var err error
	var uid uint64

	// TODO：对验证码进行校验


	rpcGetFromPhone, err := UserPbClient.GetFromPhone(context.Background(), &pbUser.GetFromPhoneReq{
		Phone: req.Phone,
	})
	if err != nil {
		logging.Logger().Error(err)
		xhttp.SendJsonResponse(c, err, rsp)
		return
	}
	if !rpcGetFromPhone.BaseResponse.Success{
		logging.Logger().Error(rpcGetFromPhone.BaseResponse.Error)
		xhttp.SendJsonResponse(c, errno.ErrUserLogin, rsp)
		return
	}


	if nil != rpcGetFromPhone.UserInfo{
		uid = rpcGetFromPhone.UserInfo.Uid
	}

	// 判断用户是否存在
	if nil == rpcGetFromPhone.UserInfo {
		// TODO: 用户不存在进行注册
		rpcAddUser,err  := UserPbClient.AddUser(context.Background(), &pbUser.AddUserReq{
			Phone: req.Phone,
			Nick:  "用户",
		})
		if err !=nil{
			logging.Logger().Error(err)
			xhttp.SendJsonResponse(c, err, rsp)
			return
		}
		if !rpcAddUser.BaseResponse.Success{
			logging.Logger().Error(rpcAddUser.BaseResponse.Error)
			xhttp.SendJsonResponse(c, errno.ErrUserLogin, rsp)
			return
		}


		uid = rpcAddUser.Uid
	}

	// 生成token
	rpcGenerateToken, err  := UserPbClient.GenerateToken(context.Background(), &pbUser.GenerateTokenReq{
		Uid: uid,
	})
	if err != nil {
		logging.Logger().Error(err)
		xhttp.SendJsonResponse(c, err, rsp)
		return
	}
	if !rpcGenerateToken.BaseResponse.Success{
		logging.Logger().Error(rpcGenerateToken.BaseResponse.Error)
		xhttp.SendJsonResponse(c, errno.ErrUserLogin, rsp)
		return
	}


	rsp.Token = rpcGenerateToken.Token
	xhttp.SendJsonResponse(c, err, rsp)
}

type PhoneLoginReq struct {
	Phone string `json:"phone" form:"phone"` // 手机号
	Code  string `json:"code" form:"code"`   // 验证码
}

type PhoneLoginRsp struct {
	Token string `json:"string"`
}
