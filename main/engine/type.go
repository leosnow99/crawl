package engine

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult
}

func NilParse(data []byte) ParseResult {
	return ParseResult{}
}

type Schedule interface {
	Submit(Request)
	Run()
	WorkReady(chan Request)
	WorkChan() chan Request
}
