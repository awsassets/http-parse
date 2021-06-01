package parser

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

type Url struct {
	Protocol string
	Host     string
	Port     int
	Path     string
	Params   map[string]string
}
