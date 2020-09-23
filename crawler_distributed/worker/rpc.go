package worker

import "github.com/zhibailin/go-distributed-crawler-from-scratch/crawler/engine"

type CrawlService struct {
}

func (CrawlService) Process(req Request, result *ParseResult) error {
	// 作为 Service，接收的 Request 是序列化后的，需要反序列化处理，才能传给 engine.Worker 处理
	engineReq, err := DeserializeRequest(req)
	if err != nil {
		return err
	}
	engineResult, err := engine.Worker(engineReq)
	if err != nil {
		return nil
	}

	// Worker 处理后的结果，需要序列化才能发出去
	*result = SerializeResult(engineResult)
	return nil // 注意指针传递下，只用返回 Type error
}
