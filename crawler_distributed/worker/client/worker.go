package client

import (
	"log"
	"net/rpc"

	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler_distributed/rpcsupport"

	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler/engine"
	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler_distributed/config"
	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler_distributed/worker"
)

// 不在这里用 NewClient 直接创建一个 client，
// 而是从然后拿出一个传入
// 用 channel 传，而不是函数参数，因为 pool 启用 goroutine 并发分发
func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	return func(req engine.Request) (engine.ParseResult, error) {
		// engine 发 request 给 worker，先将 request 序列化
		sReq := worker.SerializeRequest(req)
		var sResult worker.ParseResult
		client := <-clientChan
		err := client.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		// worker 返回 result 给 engine，需要反序列化
		return worker.DeserializeResult(sResult), nil
	}
}

func CreateClientPool(hosts []string) chan *rpc.Client {
	// 根据 hosts 创建 len(hosts) 个 clients 构成 pool
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("connected to %s", h)
		} else {
			log.Printf("error connecting to %s: %v", h, err)
		}
	}

	// 采用嵌套 for，因为不止发一轮 clients
	// 上面创建 clients 的 `for range` 和下面的分发 client 的 `for range` 是并发执行的
	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
