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

// Query 查询指定日期的日程列表
// @Summary      查询日程
// @Description  根据用户ID和日期查询日程列表
// @Tags         日程管理
// @Accept       json
// @Produce      json
// @Param        request  body      schedule.QueryReq  true  "查询参数"
// @Success      200      {object}  system.Response{data=[]schedule.Schedule}
// @Failure      500      {object}  system.Response
// @Router       /schedule/query [post]
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

// QueryMonth 查询指定月份的日程列表
// @Summary      查询月度日程
// @Description  根据用户ID和年月查询整月的日程列表
// @Tags         日程管理
// @Accept       json
// @Produce      json
// @Param        request  body      schedule.QueryReq  true  "查询参数"
// @Success      200      {object}  system.Response{data=[]schedule.Schedule}
// @Failure      500      {object}  system.Response
// @Router       /schedule/queryMonth [post]
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

// Store 创建新日程
// @Summary      创建日程
// @Description  创建一个新的日程安排
// @Tags         日程管理
// @Accept       json
// @Produce      json
// @Param        request  body      schedule.StoreReq  true  "日程信息"
// @Success      200      {object}  system.Response
// @Failure      500      {object}  system.Response
// @Router       /schedule/store [post]
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

// Update 更新日程
// @Summary      更新日程
// @Description  更新已存在的日程信息
// @Tags         日程管理
// @Accept       json
// @Produce      json
// @Param        request  body      schedule.UpdateReq  true  "更新信息"
// @Success      200      {object}  system.Response
// @Failure      500      {object}  system.Response
// @Router       /schedule/update [post]
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
