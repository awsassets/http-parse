package parser

import (
	"encoding/json"
	"github.com/EmYiQing/http-parser/util"
	"io/ioutil"
	"strings"
)

func FromReqFile(filename string) (req Request) {
	requestByte, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	request := string(requestByte)
	return doParseRequest(request)
}

func FromReqString(request string) (req Request) {
	return doParseRequest(request)
}

func doParseRequest(request string) (req Request) {
	req = Request{}

	lineSep := util.GetLineSep()

	dataTemp := strings.Split(request, lineSep+lineSep)
	if len(dataTemp) > 1 {
		req.OriginData = strings.TrimSpace(dataTemp[1])
	}
	tempReq := strings.TrimSpace(dataTemp[0])

	temp := strings.Split(tempReq, lineSep)
	if len(temp) < 1 {
		return
	}
	firstLine := temp[0]
	firstTemp := strings.Split(firstLine, " ")
	if len(firstTemp) < 3 {
		return
	}
	requestMethod := firstTemp[0]
	path := firstTemp[1]
	req.Method = requestMethod
	req.Path = path

	cookieIndex := -1
	headers := make(map[string]string)
	for i := 1; i < len(temp); i++ {
		if strings.TrimSpace(temp[i]) == "" {
			break
		}
		key := strings.Split(temp[i], ": ")[0]
		value := strings.Split(temp[i], ": ")[1]
		if key == "Cookie" {
			cookieIndex = i
			continue
		}
		headers[key] = value
	}
	req.Headers = headers

	cookies := make(map[string]string)
	if cookieIndex != -1 {
		tempCookie := strings.Split(temp[cookieIndex], ": ")[1]
		if !strings.Contains(tempCookie, "; ") {
			key := strings.Split(tempCookie, "=")[0]
			value := strings.Split(tempCookie, "=")[1]
			cookies[key] = value
		} else {
			for _, v := range strings.Split(tempCookie, "; ") {
				key := strings.Split(v, "=")[0]
				value := strings.Split(v, "=")[1]
				cookies[key] = value
			}
		}
	}
	req.Cookie = cookies

	finalData := make(map[string]string)
	if req.OriginData != "" {
		if strings.Contains(req.Headers["Content-Type"], "x-www-form") {
			items := strings.Split(req.OriginData, "&")
			for _, v := range items {
				innerTemp := strings.Split(v, "=")
				if len(innerTemp) > 1 {
					finalData[innerTemp[0]] = innerTemp[1]
				}
			}
			req.Data = finalData
		}
		if strings.Contains(req.Headers["Content-Type"], "json") {
			var result interface{}
			_ = json.Unmarshal([]byte(req.OriginData), result)
			req.Data = result
		}
	}
	return
}
