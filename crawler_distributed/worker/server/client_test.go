package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler_distributed/config"

	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler_distributed/rpcsupport"
	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler_distributed/worker"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go func() {
		_ = rpcsupport.NewServeRpc(host, worker.CrawlService{})
	}()
	time.Sleep(10 * time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url: "http://album.zhenai.com/u/1517422783",
		Parser: worker.SerializedParser{
			FunctionName: config.ParseProfile,
			Args:         map[string]string{"userId": "1517422783", "userName": "秋之回忆"},
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServieRpc, req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
