package schedule

import "crawl/main/engine"

type QueueSchedule struct {
	requestChan chan engine.Request
	workChan    chan chan engine.Request
}

func (q *QueueSchedule) Submit(request engine.Request) {
	panic("implement me")
}

func (q *QueueSchedule) ConfigureWorkChan(requests chan engine.Request) {
	q.requestChan = requests
}

func (q *QueueSchedule) WorkReady(w chan engine.Request) {
	q.workChan <- w
}

func (q *QueueSchedule) Run() {
	q.workChan = make(chan chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWork chan engine.Request

			if len(requestQ) > 0 && len(workQ) > 0 {
				activeRequest = requestQ[0]
				activeWork = workQ[0]
			}

			select {
			case r := <-q.requestChan:
				requestQ = append(requestQ, r)
			case w := <-q.workChan:
				workQ = append(workQ, w)
			case activeWork <- activeRequest:
				workQ = workQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
