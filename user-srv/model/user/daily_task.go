/*
 * @Author       : jianyao
 * @Date         : 2020-07-14 02:05:17
 * @LastEditTime : 2020-07-14 10:22:42
 * @Description  : 每日任务
 */

package user

import "time"

type DailyTask struct {
	Id         uint64    `gorm:"primary_key;comment:'主键';not null" json:"id"`
	Logo       string    `gorm:"comment:'图标'"`
	Remake     string    `gorm:"comment:'备注'"`
	Num        int64     `gorm:"comment:'累计数量'"`
	Status     bool      `gorm:"comment:'状态:true=上架,false=下架'"`
	CreateTime time.Time `gorm:"column:create_time"` // 充值时间
}

// TableName  每日任务
func (u *DailyTask) TableName() string {
	return "daily_task"
}


