package main

import (
	"crawl/main/engine"
	"crawl/main/parse"
	"crawl/main/schedule"
)

func main() {
	request := engine.Request{
		Url:       "https://book.douban.com",
		ParseFunc: parse.Content,
	}

	concurrentEngine := engine.ConcurrentEngine{
		WorkCount: 30,
		Schedule:  &schedule.QueueSchedule{},
	}
	concurrentEngine.Run(request)

	//new(engine.SimpleEngine).Run(request)

}
