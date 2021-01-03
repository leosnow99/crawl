package engine

import (
	"crawl/main/fetcher"
	"fmt"
	"log"
)

type SimpleEngine struct{}

func (s SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	
	for _, e := range seeds {
		requests = append(requests, e)
	}
	
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetch url: %s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetch error: %s\n", r.Url)
		}
		
		parseResult := r.ParseFunc(body)
		requests = append(requests, parseResult.Requests...)
		
		for _, item := range parseResult.Items {
			fmt.Printf("Get item: %s\n", item)
		}
	}
}
