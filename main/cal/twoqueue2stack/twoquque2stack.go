package twoqueue2stack


type Queue2Stack struct {
	main *Queue
	help *Queue
}

func NewStack() *Queue2Stack {
	return &Queue2Stack{
		main: New(),
		help: New(),
	}
}

func (stack *Queue2Stack) put(data interface{}) {
	stack.main.Put(data)
}


func (stack *Queue2Stack) pop() interface{}{
	if stack.main.empty() {
		return nil
	}

	for stack.main.GetSize() > 1 {
		stack.help.Put(stack.main.Pop())
	}

	result := stack.main.Pop()

	stack.main, stack.help = stack.help, stack.main
	return result
}


func (stack *Queue2Stack) putNew(data interface{}) {
	stack.main.Put(data)
}

func (stack *Queue2Stack) popNew() interface{}{
	if stack.main.empty() {
		return nil
	}
	var ret interface{}
	for !stack.main.empty() {
		ret = stack.main.Pop()
		if stack.main.empty() {
			break
		}
		stack.help.Put(ret)
	}

	stack.main, stack.help = stack.help, stack.main
	return ret
}
