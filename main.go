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
