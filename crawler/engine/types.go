package engine

type ParseFunc func(contents []byte, url string) ParseResult // // ParseFunc 是公共函数，url 可能被所有的 parser 用到，提出来

type Parser interface {
	Parse(contents []byte, url string) ParseResult // 正常解析 url，返回解析结果
	Serialized() (name string, args interface{})   // 序列化，返回函数名和函数参数
}

type Request struct {
	Url    string
	Parser Parser
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

type NilParser struct{}

func (n NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (n NilParser) Serialized() (name string, args interface{}) {
	return "NilParser", nil
}

type FuncParser struct {
	parser ParseFunc
	name   string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialized() (name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(p ParseFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}
