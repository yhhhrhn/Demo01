package service

import (
	"awesomeProject/entity"
	"awesomeProject/utils"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
	"log"
	"runtime"
	"sync"
	"time"
)

var GOMAXPROCS int = 2

func GetTaskById(id string) entity.TaskDetail {

	var taskDetail entity.TaskDetail
	var taskLog entity.TaskLog
	if len(id) > 0 {
		mysql := utils.GetMysqlDB()

		mysql.First(&taskDetail, "id=?", id)

		taskLog.Description = "gRPC查詢:" + taskDetail.Description
		taskLog.DateTime = string(time.Now().Format("2006-01-02 15:04:05"))
		mysql.Save(&taskLog)

		if taskDetail.Id != "" {
			//jsonData, _ := json.Marshal(taskDetail)
			return taskDetail
			//fmt.Println(jsonData)
		} else {
			log.Println("gRPC 查無此任務ID", id)

		}

	}
	return taskDetail
}
func GetTask(ctx *fasthttp.RequestCtx) string {
	//fmt.Println("getTask")
	var task entity.TaskDetail
	var taskLog entity.TaskLog

	id := ctx.UserValue("id")
	//if id != nil {
	mysql := utils.GetMysqlDB()
	defer mysql.Close()
	mysql.Where("id = ?", id).Find(&task)
	if task.Id != "" {
		jsonData, _ := json.Marshal(task)
		log.Println("查詢任務成功:", string(jsonData))
		taskLog.Description = "查詢任務成功:" + task.Description
		taskLog.DateTime = string(time.Now().Format("2006-01-02 15:04:05"))
		mysql.Save(&taskLog)
		return string(jsonData)
	} else {
		log.Println("查無此任務ID")
		return string("查無此任務ID")
	}

	return "查無此任務ID"
}
func ReciveTask(ctx *fasthttp.RequestCtx) {
	uuidV4 := uuid.New().String()
	jsonData := []byte(ctx.PostBody())
	var task entity.TaskDetail
	task.Id = uuidV4
	task.Status = "Pending"
	jsonErr := json.Unmarshal(jsonData, &task)
	if jsonErr != nil {
		fmt.Fprintf(ctx, jsonErr.Error())
		//fmt.Fprintf(ctx, "200")
		return
	}
	log.Println("任務內容:", task, "寫進隊列")
	entity.TaskChan <- task
	fmt.Fprintf(ctx, "200")
}

func AddTask(taskDetail entity.TaskDetail) {
	//檢查鎖
	if GetLock(taskDetail.Id) {
		//golangDateTime := time.Now().Format("2006-01-02 15:04:05")
		taskDetail.Status = "Completed"
		taskDetail.DateTime = string(time.Now().Format("2006-01-02 15:04:05"))
		var taskLog entity.TaskLog
		mysql := utils.GetMysqlDB()
		mysql.Save(&taskDetail)

		taskLog.Description = "Worker完成任務" + taskDetail.Description
		taskLog.DateTime = taskDetail.DateTime
		//記錄執行任務
		mysql.Save(&taskLog)

		//釋放鎖
		UnLock(taskDetail.Id)
		log.Println("Worker完成任務:", taskDetail)

	} else {
		///獲取鎖失敗就break
		log.Println("Worker獲取鎖失敗break:", taskDetail)

	}
	//fmt.Fprintf(ctx, "200")

}

func TaskScheduler() {
	var wg sync.WaitGroup
	//限制併發量
	runtime.GOMAXPROCS(GOMAXPROCS)
	for i := range entity.TaskChan {
		log.Println("TaskScheduler取出隊列:", i)
		wg.Add(1)
		go taskWorker(i, &wg)
	}
	wg.Wait()
	fmt.Println("所有任務執行完成")

}

func taskWorker(task entity.TaskDetail, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("Worker取得任務:", task)
	AddTask(task)

}
