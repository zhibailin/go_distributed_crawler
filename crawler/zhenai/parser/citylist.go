package parser

import (
	"fmt"
	"regexp"

	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		// m[0] 是匹配的包含m[1]、m[2] 的整个字符串
		fmt.Printf("City: %s, URL: %s\n", m[2], m[1])
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			ParseFunc: engine.NilParser,
		})
	}
	return result
}
