package ipinfo

import (
	"DcCheckerStatus/consts"
	"encoding/json"
	"os"
)

func (context *Context) doBackup(isMerge bool) {
	if isMerge {
		r, _ := json.Marshal(context.MergedDCStatus)
		_ = os.WriteFile(consts.CacheDCFile, r, 0644)
	} else {
		r, _ := json.Marshal(context.ListKnowIP)
		_ = os.WriteFile(consts.ApiFile, r, 0644)
	}
}
