package dao

import (
	"go-film-demo/model/schedule"
	"go-film-demo/plugin/db"
	"log"
	"time"

	"gorm.io/gorm"
)

// ScheduleDao 日程数据访问对象
type ScheduleDao struct {
}

// NewScheduleDao 创建日程DAO实例
func NewScheduleDao() *ScheduleDao {
	return &ScheduleDao{}
}

// ScheduleRequestVo 日程查询请求参数
type ScheduleRequestVo struct {
	UserID    int64
	Year      int16
	Month     int8
	Day       int8
	BeginTime time.Time
	EndTime   time.Time
	Content   string
	Priority  int8
	Status    int
	Paging    PageInfo // 假设有分页结构体
}

// PageInfo 分页信息
type PageInfo struct {
	Current  int // 当前页
	PageSize int // 每页大小
}

// GetPage 分页处理（假设已有实现）
func GetPage(query *gorm.DB, paging PageInfo) {
	// 这里应该是原有的分页逻辑
}

// ScheduleList 获取日程列表
func (dao *ScheduleDao) ScheduleList(vo ScheduleRequestVo) []schedule.Schedule {
	// 构建查询条件
	qw := db.Mdb.Model(&schedule.Schedule{})

	// 用户ID查询
	if vo.UserID > 0 {
		qw.Where("user_id = ?", vo.UserID)
	}

	// 年月日查询
	if vo.Year > 0 {
		qw.Where("year = ?", vo.Year)
	}
	if vo.Month > 0 {
		qw.Where("month = ?", vo.Month)
	}
	if vo.Day > 0 {
		qw.Where("day = ?", vo.Day)
	}

	// 时间范围查询
	if !vo.BeginTime.IsZero() && !vo.EndTime.IsZero() {
		qw.Where("start_time BETWEEN ? AND ?", vo.BeginTime, vo.EndTime)
	}

	// 内容模糊查询
	if vo.Content != "" {
		qw.Where("content LIKE ?", "%"+vo.Content+"%")
	}

	// 优先级查询
	if vo.Priority > 0 {
		qw.Where("priority = ?", vo.Priority)
	}

	// 状态查询
	if vo.Status > 0 {
		qw.Where("status = ?", vo.Status)
	}

	// 获取分页数据
	GetPage(qw, vo.Paging)

	// 执行查询
	var list []schedule.Schedule
	if err := qw.Order("start_time ASC").Find(&list).Error; err != nil {
		log.Println(err)
		return nil
	}
	return list
}

// CreateSchedule 创建日程
func (dao *ScheduleDao) CreateSchedule(schedule *schedule.Schedule) error {
	result := db.Mdb.Create(schedule)
	if result.Error != nil {
		log.Printf("创建日程失败: %v", result.Error)
		return result.Error
	}
	log.Printf("创建日程成功, ID: %d", schedule.ID)
	return nil
}

// UpdateSchedule 更新日程
func (dao *ScheduleDao) UpdateSchedule(s *schedule.Schedule) error {
	result := db.Mdb.Save(s)
	if result.Error != nil {
		log.Printf("更新日程失败: %v", result.Error)
		return result.Error
	}
	log.Printf("更新日程成功, ID: %d", s.ID)
	return nil
}

// UpdateScheduleByFields 更新日程部分字段
func (dao *ScheduleDao) UpdateScheduleByFields(id int64, updates map[string]interface{}) error {
	result := db.Mdb.Model(&schedule.Schedule{}).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		log.Printf("更新日程字段失败: %v", result.Error)
		return result.Error
	}
	log.Printf("更新日程字段成功, ID: %d, 影响行数: %d", id, result.RowsAffected)
	return nil
}

// DeleteSchedule 删除日程（物理删除）
func (dao *ScheduleDao) DeleteSchedule(id int64) error {
	result := db.Mdb.Delete(&schedule.Schedule{}, id)
	if result.Error != nil {
		log.Printf("删除日程失败: %v", result.Error)
		return result.Error
	}
	log.Printf("删除日程成功, ID: %d", id)
	return nil
}

// SoftDeleteSchedule 软删除日程（如果表中支持deleted_at字段）
func (dao *ScheduleDao) SoftDeleteSchedule(id int64) error {
	result := db.Mdb.Model(&schedule.Schedule{}).Where("id = ?", id).Update("status", schedule.StatusEnded) // 假设用状态表示删除
	if result.Error != nil {
		log.Printf("软删除日程失败: %v", result.Error)
		return result.Error
	}
	log.Printf("软删除日程成功, ID: %d", id)
	return nil
}

// GetScheduleByID 根据ID获取日程详情
func (dao *ScheduleDao) GetScheduleByID(id int64) (*schedule.Schedule, error) {
	var schedule schedule.Schedule
	result := db.Mdb.First(&schedule, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			log.Printf("日程不存在, ID: %d", id)
			return nil, nil
		}
		log.Printf("查询日程失败: %v", result.Error)
		return nil, result.Error
	}
	return &schedule, nil
}

// GetSchedulesByUserAndDate 根据用户ID和日期获取日程
func (dao *ScheduleDao) GetSchedulesByUserAndDate(userID int64, year int16, month int8, day int8) ([]schedule.Schedule, error) {
	var schedules []schedule.Schedule
	result := db.Mdb.Where("user_id = ? AND year = ? AND month = ? AND day = ?",
		userID, year, month, day).Order("start_time ASC").Find(&schedules)
	if result.Error != nil {
		log.Printf("查询用户日程失败: %v", result.Error)
		return nil, result.Error
	}
	return schedules, nil
}

// GetSchedulesByUserAndPriority 根据用户ID和优先级获取日程
func (dao *ScheduleDao) GetSchedulesByUserAndPriority(userID int64, priority int8) ([]schedule.Schedule, error) {
	var schedules []schedule.Schedule
	result := db.Mdb.Where("user_id = ? AND priority = ?", userID, priority).
		Order("start_time ASC").Find(&schedules)
	if result.Error != nil {
		log.Printf("查询用户优先级日程失败: %v", result.Error)
		return nil, result.Error
	}
	return schedules, nil
}

// BatchUpdateScheduleStatus 批量更新日程状态
func (dao *ScheduleDao) BatchUpdateScheduleStatus(ids []int64, status int) error {
	result := db.Mdb.Model(&schedule.Schedule{}).Where("id IN ?", ids).Update("status", status)
	if result.Error != nil {
		log.Printf("批量更新日程状态失败: %v", result.Error)
		return result.Error
	}
	log.Printf("批量更新日程状态成功, 影响行数: %d", result.RowsAffected)
	return nil
}
