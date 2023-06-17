package ipinfo

import (
	"DcCheckerStatus/consts"
	"DcCheckerStatus/http"
	"DcCheckerStatus/ipinfo/types"
	"encoding/json"
	"fmt"
	"strings"
)

func (context *Context) CheckIP(ip string) {
	r := context.ListKnowIP[ip]
	if r != nil {
		r.Ip = ip
	} else {
		result := http.ExecuteHttpRequest(fmt.Sprintf("https://ipinfo.io/%s?token=%s", ip, consts.IPInfoKey))
		var resultStruct types.IpInfoResult
		_ = json.Unmarshal([]byte(*result), &resultStruct)
		context.ListKnowIP[ip] = &types.ServerInfo{
			Ip:      ip,
			Country: resultStruct.Country,
			ASN:     strings.Split(resultStruct.Org, " ")[0],
		}
	}
}
