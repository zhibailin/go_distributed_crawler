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
	return r.ParseFunc(body), nil
}
