package persist

import (
	"context"
	"log"

	"gopkg.in/olivere/elastic.v5"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item "+"#%d: %v", itemCount, item)
			itemCount++
			_, err := Save(item)
			if err != nil {
				log.Printf("Item Saver: error "+"saving item %v: %v", item, err)
			}
		}
	}()
	return out // TODO: 返回出去的 out 和 go func() 里的 out 是什么关系
}

func Save(item interface{}) (id string, err error) {
	client, err := elastic.NewClient( // 默认到 9200 端口找服务器
		elastic.SetSniff(false)) // Must turn off in docker；因为集群不在本机，而在docker里，image不通外网，无法sniff
	if err != nil {
		return "", nil
	}

	resp, err := client.Index().Index("data_profile").Type("zhenai").BodyJson(item).Do(context.Background())

	if err != nil {
		return "", err
	}

	return resp.Id, nil
}
