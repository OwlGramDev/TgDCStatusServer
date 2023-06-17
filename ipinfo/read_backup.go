package ipinfo

import (
	"DcCheckerStatus/consts"
	"DcCheckerStatus/ipinfo/types"
	"encoding/json"
	"os"
)

func (context *Context) readBackup(isMerge bool) {
	if isMerge {
		r, err := os.ReadFile(consts.CacheDCFile)
		if err == nil {
			var recovery []types.DCStatus
			_ = json.Unmarshal(r, &recovery)
			context.MergedDCStatus = recovery
		}
	} else {
		r, err := os.ReadFile(consts.ApiFile)
		if err == nil {
			var recovery map[string]*types.ServerInfo
			_ = json.Unmarshal(r, &recovery)
			context.ListKnowIP = recovery
		}
	}
}
