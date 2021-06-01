package test

import (
	parser "github.com/EmYiQing/http-parser"
	"testing"
)

func TestReq(t *testing.T) {
	res := parser.FromReqFile("request.txt")
	t.Log("Method:", res.Method)
	t.Log("Path:", res.Path)
	t.Log("Cookies:", res.Cookie)
	t.Log("Headers:", res.Headers)
	t.Log("OriginData:", res.OriginData)
	t.Log("Data", res.Data)
}

func TestResp(t *testing.T) {
	res := parser.FromRespFile("response.txt")
	t.Log("Version", res.Version)
	t.Log("Code", res.Code)
	t.Log("Reason:", res.Reason)
	t.Log("Headers:", res.Headers)
	t.Log("Body:", res.Body)
}

func TestUrl(t *testing.T) {
	originUrl := "https://www.xxx.com/index.php?id=1&username=admin"
	res := parser.ResolveUrl(originUrl)
	t.Log("Protocol:", res.Protocol)
	t.Log("Host:", res.Host)
	t.Log("Port", res.Port)
	t.Log("Path:", res.Path)
	t.Log("Params:", res.Params)
}
