package user

import (
	"fmt"
	"log"
	"sync"

	"micro_demo/basic/db/xgorm"
	proto "micro_demo/proto/user"
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
	// QueryUserByName 根据用户名获取用户
	QueryUserByName(userName string) (ret *proto.User, err error)
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

}

// User ...
type User struct {
	ID       int64  `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	UserID   int64  `json:"user_id" gorm:"column:user_id"`
	UserName string `json:"user_name"  gorm:"column:user_id"`
	Pwd      string `json:"pwd"  gorm:"column:pwd"`
}

// QueryUserByName 查询名字
// func (s *service) QueryUserByName(userName string) (ret *proto.User, err error) {
// 	queryString := `SELECT user_id, user_name, pwd FROM user WHERE user_name = ?`

// 	// 获取数据库
// 	o := db.GetDB()

// 	ret = &proto.User{}

// 	// 查询
// 	err = o.QueryRow(queryString, userName).Scan(&ret.Id, &ret.Name, &ret.Pwd)
// 	if err != nil {
// 		log.Infof("[QueryUserByName] 查询数据失败，err：%s", err)
// 		return
// 	}
// 	return
// }

func (s *service) QueryUserByName(userName string) (ret *proto.User, err error) {
	data := User{}
	ret = &proto.User{}
	o := xgorm.GetDB()

	err = o.Where("user_name = ?", userName).First(&data).Error

	if err != nil {
		log.Fatal(err)
		return ret, err
	}

	ret = &proto.User{
		Pwd: data.Pwd,
	}
	return ret, nil
}
