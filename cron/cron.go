package cron

import (
	"awesomeProject/service"
	"github.com/robfig/cron/v3"
	"log"
	"time"
)

type TaskJob struct {
}

func (job TaskJob) Run() {
	log.Println("Run 60s TaskJob")
	service.TaskScheduler()
}

func InitCron() {
	// 创建一个默认的cron对象
	c := cron.New()

	// 添加任务
	//c.AddFunc("@every 15s", func() {
	//	log.Printf("15s Scheduler")
	//	service.TaskScheduler()
	//})
	spec := "@every 60s"
	c.AddJob(spec, TaskJob{})
	//开始执行任务
	c.Start()

	//defer c.Stop()

	time.Sleep(time.Second * 5)
	//阻塞
	//select {}
}
