package persist

import (
	"context"
	"errors"
	"log"

	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler/engine"

	"gopkg.in/olivere/elastic.v5"
)

func ItemSaver() chan engine.Item {
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item "+"#%d: %v", itemCount, item)
			itemCount++
			err := Save(item)
			if err != nil {
				log.Printf("Item Saver: error "+"saving item %v: %v", item, err)
			}
		}
	}()
	return out // TODO: 返回出去的 out 和 go func() 里的 out 是什么关系
}

func Save(item engine.Item) error {
	client, err := elastic.NewClient( // 默认到 9200 端口找服务器
		elastic.SetSniff(false)) // Must turn off in docker；因为集群不在本机，而在docker里，image不通外网，无法sniff
	if err != nil {
		return err
	}
	// 处理 Type 为空
	if item.Type == "" {
		return errors.New("must supply Type")
	}
	// 处理 Id 为空：先采用 Elasticsearch 自动填充的 Id(string)，若 item.Id 不为空再置换
	indexService := client.Index().Index("data_profile").Type(item.Type).BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err = indexService.Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}
