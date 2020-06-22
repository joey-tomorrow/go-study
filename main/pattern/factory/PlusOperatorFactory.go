package factory

type PlusOperatorFactory struct {
}

type PlusOperator struct {
	*OperatorBase
}

func (this *PlusOperator) SetLeft(i int) {
	this.left = i
}

func (this *PlusOperator) SetRight(i int) {
	this.right = i
}

func (this *PlusOperator) Result() int {
	return this.left + this.right
}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		&OperatorBase{},
	}
}
