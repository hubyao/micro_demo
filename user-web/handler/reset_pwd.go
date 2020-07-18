package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"micro_demo/comm/logging"
	"micro_demo/comm/xhttp"
	"micro_demo/comm/xhttp/errno"
	pbUser "micro_demo/proto/user"
)

func RestPwd(c *gin.Context) {
	req := restPwdReq{}
	rsp := restPwdRsp{}
	var err error
	var uid uint64

	// 对验证码进行校验
	rpcVerifyCodeSms, err := UserPbClient.VerifyCodeSms(context.Background(), &pbUser.VerifyCodeSmsReq{
		Phone:   req.Phone,
		Code:    req.Code,
		SmsType: "login",
	})
	if err != nil {
		logging.Logger().Error(err)
		xhttp.FailRsp(c, errno.ErrSmsCodeInvalid, err.Error())
		return
	}
	if !rpcVerifyCodeSms.BaseResponse.Success {
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
	if !rpcGetFromPhone.BaseResponse.Success {
		logging.Logger().Error(rpcGetFromPhone.BaseResponse.Error)
		xhttp.FailRsp(c, errno.ErrUserLogin, "")
		return
	}

	if nil != rpcGetFromPhone.UserInfo {
		uid = rpcGetFromPhone.UserInfo.Uid
	}

	// 用户不已存在
	if 0 != rpcGetFromPhone.UserInfo.Uid {
		xhttp.FailRsp(c, errno.ErrUserExist, "")
		return
	}

	// 添加密码
	_, err = UserPbClient.UpdatePwd(context.Background(), &pbUser.UpdatePwdReq{
		Uid: uid,
		Pwd: req.Pwd,
	})

	if err != nil {
		logging.Logger().Error(err)
		xhttp.FailRsp(c, err, "")
		return
	}

	xhttp.OkRsp(c, rsp)
}

type restPwdReq struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
	Pwd   string `json:"pwd"`
}

type restPwdRsp struct {
}
