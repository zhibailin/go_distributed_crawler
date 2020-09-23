package worker

import "github.com/zhibailin/go-distributed-crawler-from-scratch/crawler/engine"

type CrawlService struct {
}

func (CrawlService) Process(req engine.Request, result *engine.ParseResult) error {
	// TODO
}
