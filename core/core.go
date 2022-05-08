package core

import (
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	Version = "0.1"
	AppName = "go-proxy"
	Desc    = "网络代理"
)

func SimpleResponse(status int) *http.Response {
	switch status {
	case http.StatusOK:
		return &http.Response{StatusCode: http.StatusOK}
	case http.StatusBadGateway:
		fallthrough
	default:
		resp := new(http.Response)
		resp.StatusCode = http.StatusBadGateway
		resp.Header = make(http.Header)
		resp.Header.Set("Content-Type", "text/html; charset=utf-8")
		resp.Body = ioutil.NopCloser(strings.NewReader("转发出错"))
		return resp
	}
}
