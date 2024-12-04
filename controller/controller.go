package controller

import (
	"awesomeProject/service"
	"fmt"
	"github.com/valyala/fasthttp"
)

func GetTask(ctx *fasthttp.RequestCtx) {

	//validate := validator.New()

	fmt.Fprintf(ctx, service.GetTask(ctx))

}

func PostTask(ctx *fasthttp.RequestCtx) {
	service.ReciveTask(ctx)

}
