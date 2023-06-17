package handlers

import (
	"DcCheckerStatus/ipinfo"
	"DcCheckerStatus/ipinfo/types"
	"encoding/json"
	"github.com/valyala/fasthttp"
)

func SendData(ip string, ctx *fasthttp.RequestCtx, ipInfo *ipinfo.Context) {
	if ctx.IsPost() {
		ctx.SetStatusCode(fasthttp.StatusOK)
		var listStatus []types.DCStatus
		err := json.Unmarshal(ctx.PostBody(), &listStatus)
		if err == nil {
			ipInfo.CheckIP(ip)
			ipInfo.RegisterData(ip, listStatus)
			ctx.SetConnectionClose()
			return
		}
	}
	Forbidden(ctx)
}
