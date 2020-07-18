package sms

import (
	"fmt"
	"github.com/qichengzx/qcloudsms_go"
	"micro_demo/comm/logging"
)

type smsSrv struct {
	qcloudSMS *qcloudsms.QcloudSMS
}

func NewCli(appid string, appkey string, sign string) *smsSrv {
	opt := qcloudsms.NewOptions(appid, appkey, sign)
	return &smsSrv{
		qcloudSMS: qcloudsms.NewClient(opt),
	}

}

// SendSms 发送验证码
// mobile 手机号
// msg 短信消息
func (s smsSrv) SendSms(mobile string, msg string)(bool,error) {
	sm := qcloudsms.SMSSingleReq{
		Type: 0,
		Msg:  msg, // 短信内容
		Tel:  qcloudsms.SMSTel{Nationcode: "86", Mobile: mobile},
	}

	ok,err := s.qcloudSMS.SendSMSSingle(sm)
	if err != nil {
		logging.Logger().Error(err)
		return false,err
	}

	return ok,nil
}

// SendCodeSms 发送验证码
func (s smsSrv) SendCodeSms(mobile string, code string)(bool,error) {
	sm := qcloudsms.SMSSingleReq{
		Type: 0,
		Msg:  fmt.Sprintf("您的验证码为：%v。该验证码10分钟内有效。",code), // 短信内容
		Tel:  qcloudsms.SMSTel{Nationcode: "86", Mobile: mobile},
	}

	ok,err := s.qcloudSMS.SendSMSSingle(sm)
	if err != nil {
		logging.Logger().Error(err)
		return false,err
	}

	return ok,nil
}
