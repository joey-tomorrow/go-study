package linked2queue2stack

type Queue struct {
	*DoubleEndLinedQueue
}

func (s *Queue) Put(value interface{}) {
	s.putFromHead(value)
}

func (s *Queue) Pop() interface{} {
	return s.popFromTail().value
}
