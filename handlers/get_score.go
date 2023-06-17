package handlers

import (
	"DcCheckerStatus/ipinfo"
	"DcCheckerStatus/ipinfo/types"
	"encoding/json"
	"github.com/valyala/fasthttp"
)

func GetScore(ip string, ctx *fasthttp.RequestCtx, ipInfo *ipinfo.Context) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ipInfo.CheckIP(ip)
	res, _ := json.Marshal(types.ScoreResult{
		Score: ipInfo.GetScoreRate(*ipInfo.ListKnowIP[ip]),
	})
	ctx.SetBody(res)
	ctx.SetConnectionClose()
}
