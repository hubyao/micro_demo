/*
 * @Author       : jianyao
 * @Date         : 2020-07-16 09:47:38
 * @LastEditTime : 2020-07-16 10:12:05
 * @Description  :  redis_key
 *  所有rediskey定义在该文件中
 * - redis_key的格式以  xxx:xxx隔开
 *  注释 存储类型=>rediskey的注释=>存储的结构与
 */
package xredis

import (
	"time"
)

var (
	RedisKeyPhoneCodeSms = CacheKeyObj{"phone:%v:sms_type:%v",10 * time.Minute} // set => phone: 手机号:sms_type:验证码类型 => 验证码  
)

type CacheKeyObj struct {
	Key string        `json:"key"` // key
	Exp time.Duration `json:"exp"` // 过期时间
}
