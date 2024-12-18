package main

import (
	"awesomeProject/controller"
	"awesomeProject/cron"
	gRpc "awesomeProject/grpc"
	_ "awesomeProject/utils"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"log"
)

func main() {

	router := router.New()

	router.GET("/tasks/{id}", controller.GetTask)

	router.POST("/tasks", controller.PostTask)

	log.Println("Web Server is Running listening port 9299")
	cron.InitCron()
	go gRpc.InitGrpc()

	log.Fatal(fasthttp.ListenAndServe(":9299", router.Handler))

}
