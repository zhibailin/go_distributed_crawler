package engine

import (
	"log"
)

// 实现一个 ConcurrentEngine "类"
type ConcurrentEngine struct {
	Scheduler   Scheduler // "构造函数"，指定该 engine 采用的 Scheduler
	WorkerCount int       // 指定并发的 worker 数量
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	// 初始化管道配置
	out := make(chan ParseResult)
	e.Scheduler.Run()

	// engine 调用 NewWorker 并发创建 workers，
	// 并发执行 fetch 和 parse
	for i := 0; i < e.WorkerCount; i++ {
		NewWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
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
