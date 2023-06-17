package ipinfo

import (
	"DcCheckerStatus/ipinfo/types"
	"time"
)

func (context *Context) RegisterData(ip string, status []types.DCStatus) {
	context.ListKnowIP[ip].Status = &types.JSONStatus{}
	context.ListKnowIP[ip].Status.Status = status
	context.ListKnowIP[ip].Status.LastRefresh = time.Now().Unix()
}
