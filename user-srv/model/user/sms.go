/**
短信发送
*/
package user

import (
	"fmt"
	"math/rand"
	"micro_demo/basic/redis"
	"micro_demo/comm/logging"
	"micro_demo/comm/qcloud/sms"
	"micro_demo/comm/xredis"
)

// SendCodeSms 发送验证码短信
func (s *service)SendCodeSms(phone, smsType string) error {
	code := rand.Intn(99999) + 1000000 // 生成六位数验证码
	_, err := sms.NewCli("appid", "appkey", "sign").
		SendCodeSms(phone, fmt.Sprintf("%v", code))
	if err != nil {
		logging.Logger().Error(err)
		return err
	}

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
