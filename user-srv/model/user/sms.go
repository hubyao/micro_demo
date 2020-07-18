/**
短信发送
*/
package user

import (
	"fmt"
	"math/rand"
	"micro_demo/basic/redis"
	"micro_demo/comm/logging"
	"micro_demo/comm/xredis"
)

// SendCodeSms 发送验证码短信
func (s *service)SendCodeSms(phone, smsType string) (err error) {
	code := rand.Intn(99999) + 100000 // 生成六位数验证码
	logging.Logger().Debugf("验证码: %v",code)
	//_, err := sms.NewCli("appid", "appkey", "sign").
	//	SendCodeSms(phone, fmt.Sprintf("%v", code))
	//if err != nil {
	//	logging.Logger().Error(err)
	//	return err
	//}

	// 对验证码进行缓存
	cacheKey := fmt.Sprintf(xredis.RedisKeyPhoneCodeSms.Key, phone, smsType)
	err = redis.GetRedis().Set(cacheKey, code, xredis.RedisKeyPhoneCodeSms.Exp).Err()
	if err != nil {
		logging.Logger().Error(err)
	}

	return nil
}

// 验证redisKey
func (s *service)VerifyCodeSms(phone, code, smsType string) (ok bool) {
	// 对比key
	cacheKey := fmt.Sprintf(xredis.RedisKeyPhoneCodeSms.Key, phone, smsType)
	data ,_ := redis.GetRedis().Get(cacheKey).Result()
	if data != code {
		return false
	}

	// 删除key
	redis.GetRedis().Del(cacheKey)
	return true
}
