package antiflood

import (
	"DcCheckerStatus/antiflood/types"
	"DcCheckerStatus/consts"
	"fmt"
	"time"
)

func (context *Context) CheckFlood(key string) bool {
	t := time.Now().Unix()
	currentFloodWait := context.ListFloodData[key]
	if currentFloodWait.LastRequest == 0 {
		currentFloodWait = types.Data{
			LastRequest: t,
		}
	}
	if currentFloodWait.Requests >= consts.MaxRequestsPerIP {
		banTime := (currentFloodWait.LastRequest + consts.BanTime) - t
		if banTime > 0 {
			fmt.Println(key)
			return false
		} else {
			currentFloodWait.LastRequest = t
			currentFloodWait.Requests = 1
		}
	} else {
		if t-currentFloodWait.LastRequest > consts.MaxTime {
			currentFloodWait.Requests = 0
			currentFloodWait.LastRequest = t
		} else {
			currentFloodWait.Requests += 1
		}
	}
	context.ListFloodData[key] = currentFloodWait
	return true
}
