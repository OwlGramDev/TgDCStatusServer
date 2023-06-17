package main

import (
	"DcCheckerStatus/antiflood"
	"DcCheckerStatus/consts"
	"DcCheckerStatus/handlers"
	"DcCheckerStatus/http/webserver"
	"DcCheckerStatus/ipinfo"
	"github.com/valyala/fasthttp"
)

var ipInfo *ipinfo.Context
var antiFlood *antiflood.Context

func main() {
	consts.LoadEnv()
	client := webserver.Client(onMessage)
	ipInfo = ipinfo.Client()
	antiFlood = antiflood.Client()
	go ipInfo.Run()
	client.Run()
}

func onMessage(ctx *fasthttp.RequestCtx) {
	ip := string(ctx.Request.Header.Peek("X-Forwarded-For"))
	if antiFlood.CheckFlood(ip) {
		switch string(ctx.Path()) {
		case "/sendData":
			handlers.SendData(ip, ctx, ipInfo)
			break
		case "/getScore":
			handlers.GetScore(ip, ctx, ipInfo)
			break
		default:
			handlers.NotFound(ctx)
			break
		}
	} else {
		ctx.SetStatusCode(fasthttp.StatusTooManyRequests)
		ctx.SetConnectionClose()
	}
}
