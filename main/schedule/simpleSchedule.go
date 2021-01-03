package schedule

import "crawl/main/engine"

type Schedule interface {
	Submit(engine.Request)
	ConfigureWorkChan(chan engine.Request)
	Run()
	WorkReady(chan engine.Request)
}

type SimpleSchedule struct {
	workerChan chan engine.Request
}

func (s *SimpleSchedule) Run() {
	panic("implement me")
}

func (s *SimpleSchedule) WorkReady(requests chan engine.Request) {
	panic("implement me")
}

func (s *SimpleSchedule) Submit(request engine.Request) {
	go func() {
		s.workerChan <- request
	}()
}

func (s *SimpleSchedule) ConfigureWorkChan(c chan engine.Request) {
	s.workerChan = c
}
