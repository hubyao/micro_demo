package sms

import "testing"

func TestSendSms(t *testing.T) {
	ok,err := NewCli("","","").SendSms("111111111","短信消息")
	t.Log(ok,err)
}