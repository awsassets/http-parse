package parser

import (
	"regexp"
	"strconv"
	"strings"
)

func ResolveUrl(url string) Url {
	ret := Url{}
	protocolRe := regexp.MustCompile("(http.*?)://")
	protocol := protocolRe.FindAllStringSubmatch(url, -1)[0][1]
	ret.Protocol = protocol
	hostRe := regexp.MustCompile("http.*?://(.*?)/")
	host := hostRe.FindAllStringSubmatch(url, -1)[0][1]
	port := 80
	ret.Host = host
	if strings.Contains(host, ":") {
		port, _ = strconv.Atoi(strings.Split(host, ":")[1])
		ret.Host = strings.Split(host, ":")[0]
	}
	ret.Port = port
	pathRe := regexp.MustCompile("http.*?://.*?(/.*)\\?")
	path := pathRe.FindAllStringSubmatch(url, -1)[0][1]
	ret.Path = path
	paramsRe := regexp.MustCompile("http.*?://.*?\\?(.*)")
	paramsStr := paramsRe.FindAllStringSubmatch(url, -1)[0][1]
	if strings.TrimSpace(paramsStr) != "" {
		params := strings.Split(paramsStr, "&")
		paramsMap := make(map[string]string)
		for _, v := range params {
			temp := strings.Split(v, "=")
			key := temp[0]
			value := temp[1]
			paramsMap[key] = value
		}
		ret.Params = paramsMap
	}
	return ret
}
