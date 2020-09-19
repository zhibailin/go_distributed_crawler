package main

import (
	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler/engine"
	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler/persist"
	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler/scheduler"
	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{}, // 因为是指针接收者，所以加&
		WorkerCount: 100,
		ItemChan:    persist.ItemSaver(), // 注意这里有()，执行了 ItemSaver
	}
	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun/shanghai",
		ParseFunc: parser.ParseCity,
	})
}
