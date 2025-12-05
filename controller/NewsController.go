package controller

import (
	"go-film-demo/dao"
	"go-film-demo/model/news"
	"go-film-demo/model/system"
	"go-film-demo/plugin/spider"
	"time"

	"github.com/gin-gonic/gin"
)

var NewsDao = dao.NewNewsRepository()

// Start 启动新闻采集
// @Summary      启动新闻采集
// @Description  手动触发新闻数据采集任务
// @Tags         新闻管理
// @Produce      json
// @Success      200  {object}  system.Response
// @Router       /news/start [get]
func Start(c *gin.Context) {
	go spider.CollectNews()
	system.Success("ok", "已启动", c)
}

// QueryNews 查询新闻列表
// @Summary      查询新闻
// @Description  根据日期查询新闻列表
// @Tags         新闻管理
// @Accept       json
// @Produce      json
// @Param        request  body      news.QueryReq  true  "查询参数"
// @Success      200      {object}  system.Response{data=[]news.News}
// @Failure      500      {object}  system.Response
// @Router       /news/query [post]
func QueryNews(c *gin.Context) {
	var req news.QueryReq
	if err := c.ShouldBindJSON(&req); err != nil {
		system.Failed(err.Error(), c)
	}
	parse, err := time.ParseInLocation("2006-01-02 15:04:05", req.Date, time.Local)
	if err != nil {
		system.Failed(err.Error(), c)
	}
	dateRange, err := NewsDao.GetNewsByDateRange(parse)
	system.Success(dateRange, "ok", c)
}
