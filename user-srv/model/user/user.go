/*
 * @Author       : jianyao
 * @Date         : 2020-07-14 02:05:17
 * @LastEditTime : 2020-07-14 03:41:07
 * @Description  : file content
 */ 

package user

import(
	"micro_demo/comm/logging"
	"time"
	"micro_demo/basic/db/xgorm"
)

 // User ...
type User struct {
	UId        uint64    `gorm:"primary_key;column:uid;type:bigint(20);not null" json:"-"`
	Phone      string    `gorm:"unique;column:phone;type:varchar(255)" json:"phone"` // 手机号
	Nick       string    `gorm:"column:nick;type:varchar(255)" json:"nick"`          // 昵称
	Createtime time.Time `gorm:"column:createtime;type:datetime" json:"createtime"`
	Updatetime time.Time `gorm:"column:updatetime;type:datetime" json:"updatetime"` // 更新时间
}

// TableName get sql table name.获取数据库表名
func (u *User) TableName() string {
	return "user"
}


// Add  添加
func (s *service)AddUser(data *User)(err error){
	err = xgorm.GetDB().Table((&User{}).TableName()).Create(data).Error
	if err !=nil {
		logging.Logger().Error(err)	
	}

	return
}

// Update 更新
func (s *service)UpdateUser(uId uint64,data User)(err error){
	err = xgorm.GetDB().Table((&User{}).TableName()).Update(data).Error
	if err !=nil {
		logging.Logger().Error(err)	
	}

	return
}

// GetFromUId 通过uid获取用户
func (s *service) GetFromUId(uid uint64) (result *User, err error) {
	result = &User{}
	err = xgorm.GetDB().Table((&User{}).TableName()).Where("uid = ?", uid).Find(result).Error
	if err !=nil {
		logging.Logger().Error(err)	
	}

	return
}

// GetBatchFromUId 批量唯一主键查找
func (s *service) GetBatchFromUId(uids []uint64) (results []*User, err error) {
	err = xgorm.GetDB().Table((&User{}).TableName()).Where("uid IN (?)", uids).Find(&results).Error
	if err !=nil {
		logging.Logger().Error(err)	
	}


	return
}

// GetFromPhone 通过phone获取用户
func (s *service) GetFromPhone(phone string) (result *User, err error) {
	result = &User{}
	
	err = xgorm.GetDB().Table((&User{}).TableName()).Where("phone = ?", phone).Find(result).Error
	if err !=nil {
		logging.Logger().Error(err)	
	}

	return
}

// GetBatchFromPhone 批量手机号查找 
func (s *service) GetBatchFromPhone(phones []string) (results []*User, err error) {
	err = xgorm.GetDB().Table((&User{}).TableName()).Where("phone IN (?)", phones).Find(&results).Error
	if err !=nil {
		logging.Logger().Error(err)	
	}

	return
}