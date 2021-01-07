package schedule

import "crawl/main/engine"

type SimpleSchedule struct {
	workerChan chan engine.Request
}

func (s *SimpleSchedule) WorkChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleSchedule) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleSchedule) WorkReady(chan engine.Request) {
	return
}

func (s *SimpleSchedule) Submit(request engine.Request) {
	go func() {
		s.workerChan <- request
	}()
}
