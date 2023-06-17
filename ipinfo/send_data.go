package ipinfo

import (
	"DcCheckerStatus/consts"
	"DcCheckerStatus/ipinfo/types"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

func SendData(data []types.DCStatus) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(fmt.Sprintf("%s?secretKey=%s", consts.OwlGramApiEndPrivate, consts.SecretDCKey))
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")
	marshal, _ := json.Marshal(data)
	req.SetBody(marshal)
	_ = fasthttp.Do(req, resp)
}
