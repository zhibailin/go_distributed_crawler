package client

import (
	"fmt"

	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler/engine"
	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler_distributed/config"
	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler_distributed/rpcsupport"
	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler_distributed/worker"
)

func CreateProcessor() (engine.Processor, error) {

	client, err := rpcsupport.NewClient(
		fmt.Sprintf(":%d", config.WorkerPort0))
	if err != nil {
		return nil, err
	}

	return func(req engine.Request) (engine.ParseResult, error) {
		// engine 发 request 给 worker，先将 request 序列化
		sReq := worker.SerializeRequest(req)
		var sResult worker.ParseResult
		err := client.Call(config.CrawlServieRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		// worker 返回 result 给 engine，需要反序列化
		return worker.DeserializeResult(sResult), nil
	}, nil
}
