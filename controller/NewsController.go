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

func Start(c *gin.Context) {
	go spider.CollectNews()
	system.Success("ok", "已启动", c)
}

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
