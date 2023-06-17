package http

import "github.com/valyala/fasthttp"

func ExecuteHttpRequest(url string) *string {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(url)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	err := fasthttp.Do(req, resp)
	if err != nil {
		return nil
	}
	if resp.StatusCode() != fasthttp.StatusOK {
		return nil
	}
	body := string(resp.Body())
	return &body
}
