package parser

import (
	"regexp"

	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler/engine"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/([0-9]+))"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`) //不能随便添加换行符。
)

func ParseCity(contents []byte, _ string) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		userUrl := string(m[1])
		id := string(m[2])
		name := string(m[3])
		result.Requests = append(result.Requests, engine.Request{
			Url:    userUrl,
			Parser: NewProfileParser(id, name),
		})
	}
	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
	}
	return result
}
