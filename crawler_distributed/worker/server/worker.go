package main

import (
	"fmt"
	"log"

	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler_distributed/config"

	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler_distributed/rpcsupport"
	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler_distributed/worker"
)

func main() {
	log.Fatal(
		rpcsupport.NewServeRpc(
			fmt.Sprintf(":%d", config.WorkerPort0),
			worker.CrawlService{}),
	)

}
