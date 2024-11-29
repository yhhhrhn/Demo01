package service

import (
	"awesomeProject/utils"
	"context"
	"time"
)

var taskKey = "taskLock:"

func GetLock(uuid string) bool {

	ctx := context.Background()
	rdb := utils.NewClient(ctx)
	res, _ := rdb.SetNX(ctx, taskKey+uuid, uuid, time.Second*5).Result()
	rdb.Close()
	return res
}

func UnLock(uuid string) bool {
	ctx := context.Background()
	rdb := utils.NewClient(ctx)
	defer rdb.Close()

	val, _ := rdb.Get(ctx, taskKey+uuid).Result()

	if val == uuid {
		//釋放鎖
		reslut, _ := rdb.Del(ctx, taskKey+uuid).Result()
		if reslut == 1 {
			return true
		} else {
			return false
		}

	} else {
		return false
	}

}
