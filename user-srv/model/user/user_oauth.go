/*
 * @Author       : jianyao
 * @Date         : 2020-07-14 09:09:53
 * @LastEditTime : 2020-07-14 10:34:12
 * @Description  : file content
 */ 
package user

import (
	"micro_demo/basic/db/xgorm"
	"micro_demo/comm/logging"
	// "time"
)

const (
	// PlatformWechat 微信平台
	PlatformWechat = "wechat"
)

// UserOauth 用户授权表
type UserOauth struct {
	UserOauthId uint64 `gorm:"primary_key;column:user_oauth_id;type:bigint(20);not null" sql:"AUTO_INCREMENT"`
	Platform    string `gorm:"column:platform;comment:'平台:wechat=微信'" json:"platform"`
	OpenId      string
	Unionid     string
	Sex         int
	Name        string
	Avatar      string
	Sessionkey  string
	UId         uint64    `gorm:"column:uid;comment:'用户id'" json:"uid"`
	xgorm.BaseModel
}

func (uo *UserOauth) TableName() string {
	return "user_oauth"
}

// AddUserOauthr  添加
func (s *service) AddUserOauthr(data *UserOauth) (err error) {
	err = xgorm.GetDB().Table((&UserOauth{}).TableName()).Create(data).Error
	if err != nil {
		logging.Logger().Error(err)
	}

	return
}

// UpdateUserOauth 更新
func (s *service) UpdateUserOauth(data UserOauth) (err error) {
	err = xgorm.GetDB().Table((&UserOauth{}).TableName()).Update(data).Error
	if err != nil {
		logging.Logger().Error(err)
	}

	return
}

// GetUserOauthByPlatformWechat 获取微信平台的数据
func (s *service) GetUserOauthByPlatformWechat(openId string) (result *UserOauth, err error) {
	result = &UserOauth{}
	err = xgorm.GetDB().Table((&UserOauth{}).TableName()).Where(`open_id =? and platform="wechat"`,openId).Find(&result).Error
	if err != nil && err != xgorm.ErrRecordNotFound{
		logging.Logger().Error(err)
		return result,err

	}

	return result,nil
}


