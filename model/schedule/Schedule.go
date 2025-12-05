package schedule

import (
	"time"
)

// Schedule 日程表结构体
type Schedule struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement;comment:日程ID" json:"id"`
	UserID    int64     `gorm:"column:user_id;default:0;not null;comment:用户ID" json:"user_id"`
	CreateAt  time.Time `gorm:"column:create_at;default:CURRENT_TIMESTAMP;not null;comment:创建时间" json:"create_at"`
	UpdateAt  time.Time `gorm:"column:update_at;default:CURRENT_TIMESTAMP;not null;onUpdate:CURRENT_TIMESTAMP;comment:更新时间" json:"update_at"`
	Year      int16     `gorm:"column:year;default:1970;not null;comment:年" json:"year"`
	Month     int8      `gorm:"column:month;default:1;not null;comment:月(1-12)" json:"month"`
	Day       int8      `gorm:"column:day;default:1;not null;comment:日(1-31)" json:"day"`
	StartTime time.Time `gorm:"column:start_time;default:CURRENT_TIMESTAMP;not null;comment:开始时间" json:"start_time"`
	EndTime   time.Time `gorm:"column:end_time;default:CURRENT_TIMESTAMP;not null;comment:结束时间" json:"end_time"`
	Content   string    `gorm:"column:content;type:varchar(500);default:'';not null;comment:日程安排内容" json:"content"`
	Priority  int8      `gorm:"column:priority;default:0;not null;comment:优先级(0-低,1-中,2-高)" json:"priority"`
	Status    int       `gorm:"column:status;default:1;not null;comment:状态：1-未开始，2-进行中，3-已结束，4-已完成" json:"status"`
}

// TableName 设置表名
func (Schedule) TableName() string {
	return "schedule"
}

// 常量定义
const (
	PriorityLow    = 0 // 低优先级
	PriorityMedium = 1 // 中优先级
	PriorityHigh   = 2 // 高优先级

	StatusNotStarted = 1 // 未开始
	StatusInProgress = 2 // 进行中
	StatusEnded      = 3 // 已结束
	StatusCompleted  = 4 // 已完成
)
