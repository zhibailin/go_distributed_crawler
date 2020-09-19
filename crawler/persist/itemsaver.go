package persist

import (
	"log"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item "+"#%d: %v", itemCount, item)
			// TODO : Save item to Elasticsearch
			itemCount++
		}
	}()
	return out // TODO: 返回出去的 out 和 go func() 里的 out 是什么关系
}
