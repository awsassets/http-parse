package parser

import (
	"github.com/EmYiQing/http-parse/util"
	"io/ioutil"
	"strconv"
	"strings"
)

func FromRespFile(filename string) (resp Response) {
	responseByte, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	response := string(responseByte)
	return doParseResponse(response)
}

func FromRespString(response string) (resp Response) {
	return doParseResponse(response)
}

func doParseResponse(response string) (resp Response) {
	resp = Response{}

	lineSep := util.GetLineSep()

	dataTemp := strings.Split(response, lineSep+lineSep)
	if len(dataTemp) > 1 {
		resp.Body = strings.TrimSpace(dataTemp[1])
	}
	tempResp := strings.TrimSpace(dataTemp[0])

	temp := strings.Split(tempResp, lineSep)
	if len(temp) < 1 {
		return
	}
	firstLine := temp[0]
	firstTemp := strings.Split(firstLine, " ")
	if len(firstTemp) < 3 {
		return
	}
	version := firstTemp[0]
	code := firstTemp[1]
	reason := firstTemp[2]
	resp.Version = version
	resp.Code, _ = strconv.Atoi(code)
	resp.Reason = reason

	headers := make(map[string]string)
	for i := 1; i < len(temp); i++ {
		if strings.TrimSpace(temp[i]) == "" {
			break
		}
		key := strings.Split(temp[i], ": ")[0]
		value := strings.Split(temp[i], ": ")[1]
		headers[key] = value
	}
	resp.Headers = headers

	return
}
