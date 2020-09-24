package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler_distributed/rpcsupport"
	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler_distributed/worker"
)

// go run worker.go --port=9000
var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(
		rpcsupport.NewServeRpc(
			fmt.Sprintf(":%d", *port),
			worker.CrawlService{}),
	)

}
