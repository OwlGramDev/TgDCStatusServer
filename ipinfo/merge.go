package ipinfo

import (
	"math"
	"time"
)

func (context *Context) Merge() {
	t := time.Now()
	context.loadData()
	totalScoreCount := int64(0)
	for i := 0; i < len(context.MergedDCStatus); i++ {
		context.MergedDCStatus[i].Ping = 0
		context.MergedDCStatus[i].Status = 0
	}
	for ip, serverInfo := range context.ListKnowIP {
		if serverInfo.Status != nil && t.Unix()-serverInfo.Status.LastRefresh <= 120 {
			serverInfo.Ip = ip
			serverInfo.Score = context.GetScoreRate(*serverInfo)
			for i := 0; i < len(serverInfo.Status.Status); i++ {
				indexSet := context.getIndexById(serverInfo.Status.Status[i].Id)
				context.MergedDCStatus[indexSet].Status += translateStatus(serverInfo.Status.Status[i].Status) * serverInfo.Score
				context.MergedDCStatus[indexSet].Ping += serverInfo.Status.Status[i].Ping * serverInfo.Score
			}
			totalScoreCount += serverInfo.Score
		}
	}
	if totalScoreCount > 0 {
		for i := 0; i < len(context.MergedDCStatus); i++ {
			context.MergedDCStatus[i].Status = translateStatus(int64(math.Round(float64(context.MergedDCStatus[i].Status) / float64(totalScoreCount))))
			context.MergedDCStatus[i].Ping = int64(math.Round(float64(context.MergedDCStatus[i].Ping) / float64(totalScoreCount)))
			if context.MergedDCStatus[i].Ping >= 1000 {
				context.MergedDCStatus[i].Status = 0
			} else if context.MergedDCStatus[i].Ping >= 300 {
				context.MergedDCStatus[i].Status = 2
			}
			if context.MergedDCStatus[i].Status == 0 {
				context.MergedDCStatus[i].LastDown = t.Unix()
			} else if context.MergedDCStatus[i].Status == 2 {
				context.MergedDCStatus[i].LastLag = t.Unix()
			}
		}
	}
}
