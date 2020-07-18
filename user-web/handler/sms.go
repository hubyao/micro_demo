/*
 * @Author       : jianyao
 * @Date         : 2020-07-14 09:09:53
 * @LastEditTime : 2020-07-14 10:51:07
 * @Description  : file content
 */ 
package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"micro_demo/comm/logging"
	"micro_demo/comm/xhttp"
	"micro_demo/comm/xhttp/errno"
	pbUser "micro_demo/proto/user"
)



func Sms(c *gin.Context) {
	req := smsReq{}
	rsp := smsRsp{}
	var err error

	// 绑定数据
	if err = c.ShouldBind(&req); err != nil {
		xhttp.FailRsp(c, err, "")
		return
	}

	// TODO: 验证手机号

	// 发送验证码
	rpcSendCodeSms,err := UserPbClient.SendCodeSms(context.Background(),&pbUser.SendCodeSmsReq{
		Phone:   req.Phone,
		SmsType: req.SmsType,
	})
	if err != nil {
		logging.Logger().Error(err)
		xhttp.FailRsp(c, errno.ErrSendSms, err.Error())
		return
	}
	if !rpcSendCodeSms.BaseResponse.Success{
		logging.Logger().Error(rpcSendCodeSms.BaseResponse.Error)
		xhttp.FailRsp(c, errno.ErrSendSms, "")
		return
	}


	xhttp.OkRsp(c, rsp)
}

type smsReq struct {
	Phone string `json:"phone" binding:"required,numeric"` // 手机号
	SmsType  string `json:"sms_type" binding:"required"` // 短信类型:login=登陆
}

type smsRsp struct {

}