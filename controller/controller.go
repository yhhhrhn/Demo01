package controller

import (
	"awesomeProject/entity"
	"awesomeProject/service"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/valyala/fasthttp"
)

type TaskDetail struct {
	//Id          string `gorm:"type:varcahr(100) NOT NULL;primary_key;" json:"id" validate:"required,min=5,max=40"`
	Description string `validate:"required,min=1,max=100"`
	//Status      string `gorm:"type:varcahr(100)" json:"status"`
	//DateTime    string `gorm:"type:varcahr(100)" json:"dateTime"`
}

func GetTask(ctx *fasthttp.RequestCtx) {
	var MyInterface interface{}
	MyInterface = ctx.UserValue("id")
	var taskDetail entity.TaskDetail
	taskDetail.Id = MyInterface.(string)
	validate := validator.New()

	// Validate the User struct
	err := validate.Struct(taskDetail)
	if err != nil {
		// Validation failed, handle the error
		//errors := err.(validator.ValidationErrors)
		ctx.WriteString("資料驗證錯誤 請檢查傳入的參數是否正確")
		ctx.SetStatusCode(400)
		//http.Error(w, fmt.Sprintf("Validation error: %s", errors), http.StatusBadRequest)
		fmt.Println(string(ctx.Response.Body()))
	} else {
		service.GetTask(ctx)
		fmt.Println(string(ctx.Response.Body()))
	}
	//fmt.Println(taskDetail.Id)

}

func PostTask(ctx *fasthttp.RequestCtx) {
	jsonData := []byte(ctx.PostBody())
	var taskDetail TaskDetail
	jsonErr := json.Unmarshal(jsonData, &taskDetail)
	if jsonErr != nil {
		fmt.Fprintf(ctx, jsonErr.Error())
		//fmt.Fprintf(ctx, "200")
		return
	}
	validate := validator.New()

	// Validate the User struct
	err := validate.Struct(taskDetail)
	if err != nil {
		// Validation failed, handle the error
		//errors := err.(validator.ValidationErrors)
		ctx.WriteString("資料驗證錯誤 請檢查傳入的參數是否正確")
		ctx.SetStatusCode(400)
		//http.Error(w, fmt.Sprintf("Validation error: %s", errors), http.StatusBadRequest)
		fmt.Println(string(ctx.Response.Body()))
	} else {
		service.ReciveTask(ctx)
		fmt.Println(string(ctx.Response.Body()))

	}
	//fmt.Println(taskDetail)

}
