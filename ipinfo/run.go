package ipinfo

import (
	"time"
)

func (context *Context) Run() {
	go func() {
		for {
			context.doBackup(false)
			time.Sleep(time.Second * time.Duration(1))
		}
	}()
	for {
		t := time.Now()
		context.Merge()
		SendData(context.MergedDCStatus)
		context.doBackup(true)
		time.Sleep(time.Second * time.Duration(60-t.Second()))
	}
}
