package client

import (
	"log"

	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler_distributed/rpcsupport"

	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler/engine"
	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler_distributed/config"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: Got item #%d: %v", itemCount, item)
			itemCount++

			result := ""
			err = client.Call(config.ItemSaverRpc, item, &result)
			if err != nil {
				log.Printf("Item Saver: error saving item %v : %v", item, err)
			}
		}
	}()
	return out, nil
}
