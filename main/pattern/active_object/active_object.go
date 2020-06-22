package main

import (
	"fmt"
	"time"
)

type OperateType int

const (
	Incr OperateType = iota
	Decr
)

type Service struct {
	queue chan OperateType
	value int
}

func (s *Service) Decr() {
	s.queue <- Decr
}

func (s *Service) Incr() {
	s.queue <- Incr
}

func NewService() *Service {
	service := &Service{
		queue: make(chan OperateType, 10),
	}

	go service.schedule()
	return service
}

func (s *Service) schedule() {
	for ot := range s.queue {
		switch ot {
		case Decr:
			s.value--
		case Incr:
			s.value++
		}
	}
}

func main() {
	service := NewService()

	for i := 0; i < 100; i++ {
		go func() {
			service.Incr()
		}()

		go func() {
			service.Decr()
		}()
	}

	time.Sleep(2 * time.Second)
	fmt.Println(service.value)
}
