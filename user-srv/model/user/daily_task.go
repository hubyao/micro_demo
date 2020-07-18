/*
 * @Author       : jianyao
 * @Date         : 2020-07-14 02:05:17
 * @LastEditTime : 2020-07-14 10:22:42
 * @Description  : 每日任务
 */

package user

import (
	"micro_demo/basic/db/xgorm"
	"micro_demo/comm/logging"
	"time"
)

type DailyTask struct {
	DailyTaskId uint64    `gorm:"primary_key;comment:'主键';not null" json:"daily_task_id" sql:"AUTO_INCREMENT"`
	Logo        string    `gorm:"comment:'图标'"`
	Title       string    `gorm:"comment:'标题'"`
	Remake      string    `gorm:"comment:'备注'"`
	Num         int64     `gorm:"comment:'累计数量'"`
	ExpNum      int64     `gorm:"comment:'经验数'"`
	EnergyNum   int64     `gorm:"comment:'能量数'"`
	Status      bool      `gorm:"comment:'状态:true=上架,false=下架'"`
	CreateTime  time.Time `gorm:"column:create_time"`
}

// TableName  每日任务
func (u *DailyTask) TableName() string {
	return "daily_task"
}

// GetDailyTaskList 获取每日任务列表
func (s *service) GetDailyTaskList() (results []*DailyTask, err error) {
	results = make([]*DailyTask, 0)
	err = xgorm.GetDB().Table((&DailyTask{}).TableName()).
		Find(&results).Error
	if err != nil && err != xgorm.ErrRecordNotFound {
		logging.Logger().Error(err)
	}

	return
}
