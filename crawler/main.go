package main

import (
	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler/engine"
	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler/persist"
	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler/scheduler"
	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver("data_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{}, // 因为是指针接收者，所以加&
		WorkerCount: 100,
		ItemChan:    itemChan,
	}
	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
