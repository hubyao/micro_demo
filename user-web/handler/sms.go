/*
 * @Author       : jianyao
 * @Date         : 2020-07-14 09:09:53
 * @LastEditTime : 2020-07-14 10:51:07
 * @Description  : file content
 */ 
package handler

import (
	"github.com/gin-gonic/gin"
	"micro_demo/comm/xhttp"
	//pbUser "micro_demo/proto/user"
)



func Sms(c *gin.Context) {
	req := smsReq{}
	rsp := smsRsp{}
	var err error

	// 绑定数据
	if err := c.ShouldBind(&req); err != nil {
		xhttp.SendJsonResponse(c, err, nil)
		return
	}

	// TODO: 验证手机号
	// TODO: 限制发送频率30秒

	SendSms(req.Phone,req.SmsType)
	
	xhttp.SendJsonResponse(c, err, rsp)
}


func SendSms(phone string,smsType string) error  {
	return  nil
}


type smsReq struct {
	Phone string `json:"phone" binding:"required"` // 手机号
	SmsType  string `json:"sms_type" binding:"required"` // 短信类型:login=登陆
}

type smsRsp struct {

}