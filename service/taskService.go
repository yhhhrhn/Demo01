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

func GetTaskById(id string) entity.Task {

	var task entity.Task
	if len(id) > 0 {
		mysql := utils.GetMysqlDB()
		err := mysql.QueryRow("select id,description ,status,date_time from task_detail  where id=?", id).Scan(&task.Id, &task.Description, &task.Status, &task.DateTime)
		if err != nil {
			log.Println("gRPC 查無此任務ID", id)

		} else {
			jsonData, _ := json.Marshal(task)
			log.Println("gRPC 查詢任務成功:", string(jsonData))
			golangDateTime := time.Now().Format("2006-01-02 15:04:05")
			_, err := mysql.Exec("insert into task_log(description,date_time) values (?,?)", "gRPC 查詢任務:"+string(jsonData), golangDateTime)
			if err != nil {
				fmt.Printf("insert failed, err:%v\n", err)
			}
			//fmt.Fprintf(ctx, string(jsonData))
		}
	}
	return task
}
func GetTask(ctx *fasthttp.RequestCtx) {
	//fmt.Println("getTask")
	var task entity.Task

	id := ctx.UserValue("id")
	if id != nil {
		mysql := utils.GetMysqlDB()
		err := mysql.QueryRow("select id,description ,status,date_time from task_detail  where id=?", id).Scan(&task.Id, &task.Description, &task.Status, &task.DateTime)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			fmt.Fprintf(ctx, "查無此任務ID")
			log.Println("查無此任務ID", id)
			//log.Printf("data=%v", []int{1, 2, 3})

		} else {
			jsonData, _ := json.Marshal(task)
			log.Println("查詢任務成功:", string(jsonData))
			golangDateTime := time.Now().Format("2006-01-02 15:04:05")
			_, err := mysql.Exec("insert into task_log(description,date_time) values (?,?)", "查詢任務:"+string(jsonData), golangDateTime)
			if err != nil {
				fmt.Printf("insert failed, err:%v\n", err)
			}
			fmt.Fprintf(ctx, string(jsonData))
		}

	}

}
func ReciveTask(ctx *fasthttp.RequestCtx) {
	uuidV4 := uuid.New().String()
	jsonData := []byte(ctx.PostBody())
	var task entity.Task
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

func AddTask(task entity.Task) {
	//檢查鎖

	if GetLock(task.Id) {
		golangDateTime := time.Now().Format("2006-01-02 15:04:05")
		task.Status = "Completed"

		mysql := utils.GetMysqlDB()
		_, err := mysql.Exec("insert into task_detail(id,description,status,date_time) values (?,?,?,?)", task.Id, task.Description, task.Status, golangDateTime)
		if err != nil {
			fmt.Printf("insert failed, err:%v\n", err)
		}

		_, err1 := mysql.Exec("insert into task_log(description,date_time) values (?,?)", "新增任務:"+task.Description, golangDateTime)
		if err1 != nil {
			fmt.Printf("insert failed, err:%v\n", err1)
		}
		//釋放鎖
		UnLock(task.Id)
		log.Println("Worker完成任務:", task)
	} else {
		///獲取鎖失敗就break
		log.Println("Worker獲取鎖失敗break:", task)

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

func taskWorker(task entity.Task, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("Worker取得任務:", task)
	AddTask(task)

}
