package worker

import "github.com/zhibailin/go-distributed-crawler-from-scratch/crawler/engine"

type CrawlService struct {
}

func (CrawlService) Process(req engine.Request, result *engine.ParseResult) error {
	// TODO
	// engine.Request 在网络上是无法传的，因为里面的 Parser 是 interface，
	// 里面的各种函数类型无法直接传，因此需要再创建一些能传的类型，
	// 然后通过序列化函数，将 不能传的类型 转化为 能传的类型
}
