package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler_distributed/config"

	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler_distributed/rpcsupport"

	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler_distributed/persist"
	"gopkg.in/olivere/elastic.v5"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	return rpcsupport.NewServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
