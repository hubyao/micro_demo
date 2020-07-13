package user

import (
	"fmt"
	"log"
	"sync"

	"micro_demo/basic/db/xgorm"
	proto "micro_demo/proto/user"
	"micro_demo/comm/logging"
	"time"
)

var (
	s *service
	m sync.RWMutex
)

// service 服务
type service struct {
}

// Service 用户服务类
type Service interface {
	AddUser(data User)(err error) // 添加用户信息
	UpdateUser(data User)(err error) // 更新用户信息
	GetFromUId(uid uint64) (result *User, err error) // 根据uid获取用户信息
	GetBatchFromUId(uids []uint64) (results []*User, err error) // 根据uid批量获取用户信息
	GetFromPhone(phone string) (result *User, err error) // 根据手机号获取用户信息
	GetBatchFromPhone(phones []string) (results []*User, err error) // 根据手机号批量获取用户信息
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
	xgorm.GetDB().AutoMigrate(&User{})
}



// User ...
type User struct {
	UId        uint64     `gorm:"primary_key;column:uid;type:bigint(20);not null" json:"-"`
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
func (u *User)AddUser(data User)(err error){
	err = xgorm.GetDB().Table(u.TableName()).Create(data).Error
	if err !=nil {
		logging.Logger().Error(err)	
	}

	return
}

// Update 更新
func (u *User)UpdateUser(data User)(err error){
	err = xgorm.GetDB().Table(u.TableName()).Update(data).Error
	if err !=nil {
		logging.Logger().Error(err)	
	}

	return
}

// GetFromUId 通过uid获取用户
func (u *User) GetFromUId(uid uint64) (result *User, err error) {
	result = &User{}
	err = xgorm.GetDB().Table(u.TableName()).Where("uid = ?", uid).Find(result).Error
	if err !=nil {
		logging.Logger().Error(err)	
	}

	return
}

// GetBatchFromUId 批量唯一主键查找
func (u *User) GetBatchFromUId(uids []uint64) (results []*User, err error) {
	err = xgorm.GetDB().Table(u.TableName()).Where("uid IN (?)", uids).Find(&results).Error
	if err !=nil {
		logging.Logger().Error(err)	
	}


	return
}

// GetFromPhone 通过phone获取用户
func (u *User) GetFromPhone(phone string) (result *User, err error) {
	result = &User{}
	
	err = xgorm.GetDB().Table(u.TableName()).Where("phone = ?", phone).Find(result).Error
	if err !=nil {
		logging.Logger().Error(err)	
	}

	return
}

// GetBatchFromPhone 批量手机号查找 
func (u *User) GetBatchFromPhone(phones []string) (results []*User, err error) {
	err = xgorm.GetDB().Table(u.TableName()).Where("phone IN (?)", phones).Find(&results).Error
	if err !=nil {
		logging.Logger().Error(err)	
	}

	return
}