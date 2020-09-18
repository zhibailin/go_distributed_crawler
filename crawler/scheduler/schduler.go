package scheduler

import "github.com/zhibailin/go-distributed-crawler-from-scratch/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	go func() {
		s.workerChan <- request
	}()
}
