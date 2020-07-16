/*
 * @Author       : jianyao
 * @Date         : 2020-07-14 01:54:29
 * @LastEditTime : 2020-07-14 06:34:15
 * @Description  : 用户model
 */

package user

import (
	"fmt"
	//"log"
	"sync"

	"micro_demo/basic/db/xgorm"
	//proto "micro_demo/proto/user"
	"micro_demo/comm/jwt"
)

var (
	s *service
	m sync.RWMutex
)

type service struct {
}

// Service 用户服务类
type Service interface {
	// user
	AddUser(data *User) (err error)                                 // 添加用户信息
	UpdateUser(data User) (err error)                               // 更新用户信息
	GetFromUId(uid uint64) (result *User, err error)                // 根据uid获取用户信息
	GetBatchFromUId(uids []uint64) (results []*User, err error)     // 根据uid批量获取用户信息
	GetFromPhone(phone string) (result *User, err error)            // 根据手机号获取用户信息
	GetBatchFromPhone(phones []string) (results []*User, err error) // 根据手机号批量获取用户信息

	// token
	GenerateToken(userId uint64, expireDate int) (string, error) // 生成token
	ParseToken(token string) (*jwt.Claims, error)                // 解析token

	// UserOauth 用户授权
	GetUserOauthByPlatformWechat(openId string) (result *UserOauth, err error) // 获取用户授权
	UpdateUserOauth(data UserOauth) (err error)                                // 更新用户授权
	AddUserOauthr(data *UserOauth) (err error)                                 // 添加用户授权

	// 好友助力
	AddFriendHelp(data *FriendHelp) (err error)                      // 添加好友助力
	GetFriendHelpListByUser(uid uint64) (results []*User, err error) // 获取好友助力列表

	// 每日任务
	GetDailyTaskList() (results []*DailyTask, err error)


	// 验证码
	VerifyCodeSms(phone, code, smsType string) (ok bool)
	SendCodeSms(phone, smsType string) error

}

// GetService 获取服务类
func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化")
	}
	return s, nil
}

// Init 初始化用户服务层
func Init() {
	m.Lock()
	defer m.Unlock()

	if s != nil {
		return
	}

	s = &service{}

	//TODO: 同步数据库
	xgorm.GetDB().AutoMigrate(
		&User{},
		&UserOauth{},
		&FriendHelp{},
		&DailyTask{},
	)
}
