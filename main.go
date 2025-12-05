// Package main Go Schedule 日程管理系统
//
// @title           Go Schedule API
// @version         1.0
// @description     一个基于 Go + Vue3 的日程管理系统 API 文档
// @termsOfService  https://github.com/your-username/go-schedule
//
// @contact.name   API Support
// @contact.url    https://github.com/your-username/go-schedule/issues
//
// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT
//
// @host      localhost:3061
// @BasePath  /
//
// @schemes http https
package main

import (
	"fmt"
	"go-film-demo/config"
	"go-film-demo/plugin/cron"
	"go-film-demo/plugin/db"
	"go-film-demo/plugin/spider"
	"go-film-demo/router"
	"log"
)

var cronManager = cron.NewCronManager()

func init() {
	err := db.InitMysql()
	if err != nil {
		log.Fatal(err)
	}
	err = cronManager.AddEvery10SecondsTask("collect-news-task-10s", spider.CollectNews)
	if err != nil {
		log.Fatal(err)
	}
	cronManager.Start()
}

func main() {
	r := router.SetupRouter()

	err := r.Run(fmt.Sprintf(":%s", config.ListenPort))
	if err != nil {
		log.Fatal(err)
	}
}
