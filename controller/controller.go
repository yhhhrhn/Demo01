package controller

import (
	"awesomeProject/service"
	"github.com/valyala/fasthttp"
)

func GetTask(ctx *fasthttp.RequestCtx) {
	service.GetTask(ctx)
}

func PostTask(ctx *fasthttp.RequestCtx) {
	service.ReciveTask(ctx)

}
