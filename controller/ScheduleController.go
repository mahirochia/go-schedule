package controller

import (
	"fmt"
	"go-film-demo/dao"
	"go-film-demo/model/schedule"
	"go-film-demo/model/system"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var ScheduleDao = dao.NewScheduleDao()

func Query(c *gin.Context) {
	req := schedule.QueryReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		system.Failed("非法查询参数", c)
		return
	}

	vo := dao.ScheduleRequestVo{
		UserID: req.UserID,
		Year:   int16(req.Year),
		Month:  int8(req.Month),
		Day:    int8(req.Day),
	}

	scheduleList := ScheduleDao.ScheduleList(vo)
	system.Success(scheduleList, "ok", c)
}

func QueryMonth(c *gin.Context) {
	req := schedule.QueryReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		system.Failed("非法查询参数", c)
		return
	}

	vo := dao.ScheduleRequestVo{
		UserID: req.UserID,
		Year:   int16(req.Year),
		Month:  int8(req.Month),
	}

	scheduleList := ScheduleDao.ScheduleList(vo)
	system.Success(scheduleList, "ok", c)
}

func Store(c *gin.Context) {
	req := schedule.StoreReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		system.Failed("非法参数", c)
		return
	}

	start, err := stringToTimeStandard(req.Start)
	if err != nil {
		system.Failed("非法参数", c)
		return
	}

	end, err := stringToTimeStandard(req.End)
	if err != nil {
		system.Failed("非法参数", c)
		return
	}

	err = ScheduleDao.CreateSchedule(&schedule.Schedule{
		UserID:    req.UserID,
		Year:      int16(req.Year),
		Month:     int8(req.Month),
		Day:       int8(req.Day),
		StartTime: start,
		EndTime:   end,
		Content:   req.Content,
		Priority:  int8(req.Priority),
	})
	if err != nil {
		system.Failed(err.Error(), c)
		return
	}

	system.Success(nil, "ok", c)
}

func Update(c *gin.Context) {
	req := schedule.UpdateReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		system.Failed("非法参数", c)
		return
	}
	start, err := stringToTimeStandard(req.Start)
	if err != nil {
		system.Failed("非法参数", c)
		return
	}

	end, err := stringToTimeStandard(req.End)
	if err != nil {
		system.Failed("非法参数", c)
		return
	}
	s, err := ScheduleDao.GetScheduleByID(int64(req.ID))
	if err != nil {
		system.Failed("日程不存在", c)
		return
	}
	if s == nil {
		system.Failed("日程不存在", c)
		return
	}
	err = ScheduleDao.UpdateSchedule(&schedule.Schedule{
		ID:        int64(req.ID),
		UserID:    req.UserID,
		Year:      int16(req.Year),
		Month:     int8(req.Month),
		Day:       int8(req.Day),
		StartTime: start,
		EndTime:   end,
		Content:   req.Content,
		Priority:  int8(req.Priority),
		Status:    req.Status,
		CreateAt:  s.CreateAt,
		UpdateAt:  time.Now(),
	})
	if err != nil {
		system.Failed(err.Error(), c)
		return
	}

	system.Success(nil, "ok", c)

}

func stringToTimeStandard(timeStr string) (time.Time, error) {
	if strings.TrimSpace(timeStr) == "" {
		return time.Time{}, fmt.Errorf("时间字符串不能为空")
	}

	const standardFormat = "2006-01-02 15:04:05"

	t, err := time.ParseInLocation(standardFormat, timeStr, time.Local)
	if err != nil {
		return time.Time{}, fmt.Errorf("时间格式错误，期望格式: YYYY-MM-DD HH:MM:SS, 实际输入: '%s', 错误: %v", timeStr, err)
	}

	return t, nil
}
