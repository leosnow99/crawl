package engine

import (
	"crawl/main/fetcher"
	"fmt"
	"log"
)

type Schedule interface {
	Submit(Request)
	configureWorkChan(chan Request)
}

type SimpleSchedule struct {
	workerChan chan Request
}

func (s *SimpleSchedule) Submit(request Request) {
	go func() {
		s.workerChan <- request
	}()
}

func (s *SimpleSchedule) configureWorkChan(c chan Request) {
	s.workerChan = c
}

type ConcurrentEngine struct {
	WorkCount int
	Schedule  Schedule
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParseResult)
	
	e.Schedule.configureWorkChan(in)
	
	for i := 0; i < e.WorkCount; i++ {
		CreateWork(in, out)
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

func CreateWork(in chan Request, out chan ParseResult) {
	go func() {
		for {
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
