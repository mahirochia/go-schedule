package router

import (
	"go-film-demo/controller"
	"go-film-demo/plugin/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())

	schedule := r.Group("/schedule")
	{
		schedule.POST("/query", controller.Query)
		schedule.POST("/store", controller.Store)
		schedule.POST("/update", controller.Update)
		schedule.POST("/queryMonth", controller.Query)
	}

	news := r.Group("/news")
	{
		news.GET("/start", controller.Start)
		news.POST("/query", controller.QueryNews)
	}

	return r
}
