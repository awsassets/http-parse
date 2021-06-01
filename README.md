# http-parse

## Introduce

根据txt文件/字符串中的HTTP请求或响应，解析为特定的格式。常用于解析Burpsuite导出的请求/响应文件

Read the HTTP request or response in the txt file and parse it into a specific format to facilitate subsequent processing(For example, from burpsuite)

```go
type Request struct {
	Method     string
	Path       string
	Headers    map[string]string
	Cookie     map[string]string
	OriginData string
	Data       interface{}
}

type Response struct {
	Version string
	Code    int
	Reason  string
	Headers map[string]string
	Body    string
}
```

## Quick Start

```shell
go get github.com/EmYiQing/http-parse
```

```go

```