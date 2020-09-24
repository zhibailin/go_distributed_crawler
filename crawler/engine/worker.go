package engine

import (
	"log"

	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler/fetcher"
)

func Worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)

	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: errror "+"fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.Parser.Parse(body, r.Url), nil
}

func (e *ConcurrentEngine) NewWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {

	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := e.RequestProcessor(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
