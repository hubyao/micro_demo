package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"micro_demo/comm/logging"
	"micro_demo/comm/xhttp"
	"micro_demo/comm/xhttp/errno"
	pbUser "micro_demo/proto/user"
	"time"
)

func UserRegister(c *gin.Context) {
	req := userRegisterReq{}
	rsp := userRegisterRsp{}
	var err error
	//var uid uint64

	// 绑定数据
	if err := c.ShouldBind(&req); err != nil {
		xhttp.FailRsp(c, err,"")
		return
	}


	// 对验证码进行校验
	rpcVerifyCodeSms,err := UserPbClient.VerifyCodeSms(context.Background(),&pbUser.VerifyCodeSmsReq{
		Phone:   req.Phone,
		Code:    req.Code,
		SmsType: "register",
	})
	if err != nil {
		logging.Logger().Error(err)
		xhttp.FailRsp(c, errno.ErrSmsCodeInvalid, err.Error())
		return
	}
	if !rpcVerifyCodeSms.BaseResponse.Success{
		logging.Logger().Error(rpcVerifyCodeSms.BaseResponse.Error)
		xhttp.FailRsp(c, errno.ErrSmsCodeInvalid, "")
		return
	}



	// 判断用户是否存在
	rpcGetFromPhone, err := UserPbClient.GetFromPhone(context.Background(), &pbUser.GetFromPhoneReq{
		Phone: req.Phone,
	})
	if err != nil {
		logging.Logger().Error(err)
		xhttp.FailRsp(c, err, "")
		return
	}
	if !rpcGetFromPhone.BaseResponse.Success{
		logging.Logger().Error(rpcGetFromPhone.BaseResponse.Error)
		xhttp.FailRsp(c, errno.ErrUserLogin, "")
		return
	}

	logging.Logger().Debugf("debug_info %v",rpcGetFromPhone)
	//if nil != rpcGetFromPhone.UserInfo{
	//	//uid = rpcGetFromPhone.UserInfo.Uid
	//}

	// 用户已存在
	if 0 != rpcGetFromPhone.UserInfo.Uid {
		xhttp.FailRsp(c,errno.ErrUserExist,"")
		return
	}


	// 添加用户
	rpcAddUser,err  := UserPbClient.AddUser(context.Background(), &pbUser.AddUserReq{
		Phone: req.Phone,
		Nick: fmt.Sprintf("用户- %v",time.Now().Unix()) ,
	})

	if err !=nil{
		logging.Logger().Error(err)
		xhttp.FailRsp(c, err, "")
		return
	}
	if !rpcAddUser.BaseResponse.Success{
		logging.Logger().Error(rpcAddUser.BaseResponse.Error)
		xhttp.FailRsp(c, errno.ErrUserLogin, "")
		return
	}


	// 添加密码
	_,err  = UserPbClient.UpdatePwd(context.Background(), &pbUser.UpdatePwdReq{
		Uid: rpcAddUser.Uid,
		Pwd: req.Pwd,
	})

	if err !=nil{
		logging.Logger().Error(err)
		xhttp.FailRsp(c, err, "")
		return
	}

	xhttp.OkRsp(c,rsp)
}

type userRegisterReq struct {
	Phone string `json:"phone" binding:"required"`
	Code  string `json:"code" binding:"required"`
	Pwd   string `json:"pwd" binding:"required"`
}

type userRegisterRsp struct {
}
