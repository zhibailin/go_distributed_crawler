package engine

type ParseFunc func(contents []byte, url, id, name string) ParseResult

type Request struct {
	Url       string
	ParseFunc ParseFunc
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{} // 接 model.Profile
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
