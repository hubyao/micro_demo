/*
 * @Author       : jianyao
 * @Date         : 2020-07-14 02:05:17
 * @LastEditTime : 2020-07-14 10:22:42
 * @Description  : 好友助力
 */

package user

import (
	"micro_demo/basic/db/xgorm"
	"micro_demo/comm/logging"
	"time"
)

// FriendHelp
type FriendHelp struct {
	Id         uint64    `gorm:"primary_key;comment:'主键';not null" json:"id"`
	UId        uint64    `gorm:"column:uid;comment:'用户id'" json:"uid"`
	FriendUid  uint64 	 `gorm:"comment:'朋友id'"`
	CreateTime time.Time `gorm:"column:create_time"` // 充值时间
}

func (u *FriendHelp) TableName() string {
	return "friend_help"
}


// AddFriendHelp  添加 好友助力
func (s *service)AddFriendHelp(data *FriendHelp)(err error){
	err = xgorm.GetDB().Table((&FriendHelp{}).TableName()).Create(data).Error
	if err !=nil {
		logging.Logger().Error(err)
	}

	return
}

// GetFriendHelpListByuser 获取好友助力列表
func (s *service) GetFriendHelpListByuser(uid uint64) (results []*User, err error) {
	results = make([]*User,0)
	err = xgorm.GetDB().Table((&FriendHelp{}).TableName()).
		Where("uid = ?",uid).
		Joins("join user on user.uid = friend_help.uid").
		Find(&results).Error
	if err !=nil && err != xgorm.ErrRecordNotFound{
		logging.Logger().Error(err)
	}

	return
}