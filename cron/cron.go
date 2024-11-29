package cron

import (
	"awesomeProject/service"
	"github.com/robfig/cron/v3"
	"log"
	"time"
)

func InitCron() {
	// 创建一个默认的cron对象
	c := cron.New()

	// 添加任务
	c.AddFunc("@every 30s", func() {
		log.Printf("30s Scheduler")
		service.TaskScheduler()
	})

	//开始执行任务
	c.Start()
	time.Sleep(time.Second * 5)
	//阻塞
	//select {}
}
