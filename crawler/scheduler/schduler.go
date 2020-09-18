package scheduler

import "github.com/zhibailin/go-distributed-crawler-from-scratch/crawler/engine"

// 好比实现一个 Scheduler 类
type Scheduler interface {
	ConfigureWorkerChan(chan engine.Request) // 好比 Python 中的构造函数 __init__
	Submit(engine.Request)
}

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	s.workerChan <- request
}
