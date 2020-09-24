package worker

import (
	"errors"
	"fmt"
	"log"

	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler/engine"
	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler/zhenai/parser"
	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler_distributed/config"
)

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

func DeserializeRequest(r Request) (engine.Request, error) {
	p, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}

	return engine.Request{
		Url:    r.Url,
		Parser: p,
	}, nil
}

func deserializeParser(p SerializedParser) (engine.Parser, error) {
	// 方案一：将每个解析器的名字注册到一个map中
	// 方案二：用 switch ... case ...
	switch p.FunctionName {
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList), nil
	case config.ParseCity:
		return engine.NewFuncParser(parser.ParseCity, config.ParseCity), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	case config.ParseProfile:
		// 知识点：断言的用法
		// 无法直接用 p.Args["userId"]，因为 p.Args 是 interface，需要用断言获取具体的 Type map，
		// 原用 p.Args.(string)
		if args, ok := p.Args.(map[string]interface{}); ok {
			userId, idOk := args["userId"].(string)
			userName, nameOk := args["userName"].(string)
			if idOk && nameOk && (userId != "") && (userName != "") {
				return parser.NewProfileParser(userId, userName), nil
			} else {
				return nil, fmt.Errorf("invalid args: %v", p.Args)
			}
		} else {
			return nil, fmt.Errorf("type error, invalid args: %v", p.Args)
		}
	default:
		return nil, errors.New("unknown parser name")
	}
}

func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		engineReq, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserializing request: %v", err)
			continue
		}
		result.Requests = append(result.Requests, engineReq)
	}
	return result
}
