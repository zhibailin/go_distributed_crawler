package worker

import "github.com/zhibailin/go-distributed-crawler-from-scratch/crawler/engine"

// 传递函数名和函数参数，序列化出来的是
// {"ParseCityList", nil}, {"ProfileParser", userId, userName}
type SerializedParser struct {
	FunctionName string
	Args         interface{}
}

type Request struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

func SerializeRequest(r engine.Request) Request {
	functionName, args := r.Parser.Serialized()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			FunctionName: functionName,
			Args:         args,
		},
	}
}

func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}
