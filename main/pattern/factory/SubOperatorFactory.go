package factory

type SubOperatorFactory struct {
}

type SubOperator struct {
	*OperatorBase
}

func (this *SubOperator) SetLeft(i int) {
	this.left = i
}

func (this *SubOperator) SetRight(i int) {
	this.right = i
}

func (this *SubOperator) Result() int {
	return this.left - this.right
}

func (SubOperatorFactory) Create() Operator {
	return &SubOperator{
		&OperatorBase{},
	}
}
