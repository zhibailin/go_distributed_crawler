package scheduler

import "github.com/zhibailin/go-distributed-crawler-from-scratch/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerReady(requests chan engine.Request) {
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	go func() {
		s.workerChan <- request
	}()
}

type QueueScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueueScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueueScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueueScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueueScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)

	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request

			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			// 入队与出队
			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
