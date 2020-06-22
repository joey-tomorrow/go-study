package linked2queue2stack

type Queue1 struct {
	*DoubleEndLinedQueue
}

func (s *Queue1) Put(value interface{}) {
	s.putFromTail(value)
}

func (s *Queue1) Pop() interface{} {
	return s.popFromHead().value
}
