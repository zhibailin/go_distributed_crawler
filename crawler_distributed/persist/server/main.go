package main

import (
	"fmt"
	"log"

	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler_distributed/config"

	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler_distributed/rpcsupport"

	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler_distributed/persist"
	"gopkg.in/olivere/elastic.v5"
)

func main() {
	log.Fatal(serveRpc(fmt.Sprintf(":%d", config.ItemSaverPort), config.ElasticIndex))
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
