package engine

import (
	"log"

	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler/scheduler"
)

// 实现一个 ConcurrentEngine "类"
type ConcurrentEngine struct {
	Scheduler   scheduler.Scheduler // "构造函数"，指定该 engine 采用的 Scheduler
	WorkerCount int                 // 指定并发的 worker 数量
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	// 初始化管道配置
	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigureWorkerChan(in)

	// engine 调用 NewWorker 并发创建 workers，
	// 并发执行 fetch 和 parse
	for i := 0; i < e.WorkerCount; i++ {
		NewWorker(in, out)
	}

	// engine 令 Scheduler 向 管道 提交 requests
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item : %v", item)
		}
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}
