package main

import (
	"flag"
	"strings"

	workerClient "github.com/zhibailin/go-distributed-crawler-from-scratch/crawler_distributed/worker/client"

	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler/engine"
	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler/scheduler"
	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler/zhenai/parser"
	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler_distributed/config"
	itemSaverClient "github.com/zhibailin/go-distributed-crawler-from-scratch/crawler_distributed/persist/client"
)

var (
	itemSaverHost = flag.String("itermsaver_host", "", "item saver host")
	workerHosts   = flag.String("worker_hosts", "", "worker hosts(comma separated)")
)

func main() {
	flag.Parse()
	itemChan, err := itemSaverClient.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool := workerClient.CreateClientPool(strings.Split(*workerHosts, ","))
	processor := workerClient.CreateProcessor(pool)
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
