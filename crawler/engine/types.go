package engine

type ParseFunc func(contents []byte, url string) ParseResult // // ParseFunc 是公共函数，url 可能被所有的 parser 用到，提出来

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
