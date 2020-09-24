package main

import (
	"fmt"

	workerClient "github.com/zhibailin/go-distributed-crawler-from-scratch/crawler_distributed/worker/client"

	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler/engine"
	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler/scheduler"
	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler/zhenai/parser"
	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler_distributed/config"
	itemSaverClient "github.com/zhibailin/go-distributed-crawler-from-scratch/crawler_distributed/persist/client"
)

func main() {
	itemChan, err := itemSaverClient.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	processor, err := workerClient.CreateProcessor()
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueueScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})
}
