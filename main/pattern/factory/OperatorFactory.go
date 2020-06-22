package factory

type OperatorFactory interface {
	Create() Operator
}

type Operator interface {
	SetLeft(i int)
	SetRight(i int)
	Result() int
}

type OperatorBase struct {
	left, right int
}
