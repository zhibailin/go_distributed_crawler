package server

import (
	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler_distributed/rpcsupport"
	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler_distributed/worker"
)

func main() {
	// TODO
	rpcsupport.NewServeRpc(host, worker.CrawlService{})
}
