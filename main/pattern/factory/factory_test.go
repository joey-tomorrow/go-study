package factory

import (
	"fmt"
	"testing"
)

func TestPlusOperator_Result(t *testing.T) {
	var fac OperatorFactory
	fac = PlusOperatorFactory{}
	plus := fac.Create()

	plus.SetLeft(10)
	plus.SetRight(10)
	fmt.Println(plus.Result())

	sub := SubOperatorFactory{}
	subs := sub.Create()
	subs.SetLeft(10)
	subs.SetRight(10)
	fmt.Println(subs.Result())

}
