package linked2queue2stack

type Stack struct {
	*DoubleEndLinedQueue
}

func NewStack() *Stack {
	return &Stack{
		&DoubleEndLinedQueue{},
	}
}

func (s *Stack) Put(value interface{}) {
	s.putFromHead(value)
}

func (s *Stack) Pop() interface{} {
	return s.popFromHead().value
}

func (s *Stack) IsEmpty() bool {
	return s.DoubleEndLinedQueue.isEmpty()
}
