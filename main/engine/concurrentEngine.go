package engine

import (
	"crawl/main/fetcher"
	"fmt"
	"log"
)

type ConcurrentEngine struct {
	WorkCount int
	Schedule  Schedule
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)

	//调度器开始工作
	e.Schedule.Run()

	for i := 0; i < e.WorkCount; i++ {
		CreateWork(e.Schedule.WorkChan(), out, e.Schedule)
	}

	for _, request := range seeds {
		e.Schedule.Submit(request)
	}

	itemCount := 0
	for {
		resultTem := <-out
		for _, item := range resultTem.Items {
			log.Printf("Get item: %d, %s", itemCount, item)
			itemCount++
		}
		for _, requestTem := range resultTem.Requests {
			e.Schedule.Submit(requestTem)
		}
	}
}

func CreateWork(in chan Request, out chan ParseResult, s Schedule) {
	go func() {
		for {
			s.WorkReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}

			out <- result
		}
	}()
}

func worker(request Request) (ParseResult, error) {
	fmt.Println("Fetch Url: ", request.Url)

	bodyStr, err := fetcher.Fetch(request.Url)

	if err != nil {
		log.Printf("Fetch error: %s", request.Url)
		return ParseResult{}, err
	}

	return request.ParseFunc(bodyStr), nil
}
