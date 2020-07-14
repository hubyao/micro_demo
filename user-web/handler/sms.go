package handler

import (
	"github.com/gin-gonic/gin"
	"micro_demo/comm/xhttp"
	//pbUser "micro_demo/proto/user"
)



func Sms(c *gin.Context) {
	req := SmsReq{}
	rsp := SmsRsp{}
	var err error

	// TODO: 验证手机号
	// TODO: 限制发送频率30秒


	SendSms(req.Phone,req.SmsType)



	
	xhttp.SendJsonResponse(c, err, rsp)
}


func SendSms(phone string,smsType string) error  {
	return  nil
}


type SmsReq struct {
	Phone string `json:"phone" form:"phone"` // 手机号
	SmsType  string `json:"sms_type" form:"sms_type"` // 短信类型:login=登陆
}

type SmsRsp struct {

}