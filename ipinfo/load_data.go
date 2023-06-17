package ipinfo

import (
	"DcCheckerStatus/consts"
	"encoding/json"
	"os"
)

func (context *Context) loadData() {
	r, err := os.ReadFile(consts.TrustedASNFile)
	if err == nil {
		var asnList []string
		_ = json.Unmarshal(r, &asnList)
		context.TrustedASN = asnList
	}
	r, err = os.ReadFile(consts.TrustedIPFile)
	if err == nil {
		var ipList []string
		_ = json.Unmarshal(r, &ipList)
		context.TrustedIP = ipList
	}
	context.MostUsedASN = make(map[string]int64)
	for _, serverInfo := range context.ListKnowIP {
		context.MostUsedASN[serverInfo.ASN] += 1
	}
}
