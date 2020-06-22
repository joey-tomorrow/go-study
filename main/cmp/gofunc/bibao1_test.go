package gofunc

import (
	"fmt"
	"testing"
)

type Calculator struct {
	Add func(int, int) int
	Sub func(int, int) int
}

func newCalculator() Calculator {
	return Calculator{
		Add: func(a int, b int) int {
			return a + b
		},
		Sub: func(a int, b int) int {
			return a - b
		},
	}
}

func Test_bibao1(t *testing.T) {
	calc := newCalculator()
	fmt.Println(calc.Add(3, 2))
	fmt.Println(calc.Sub(3, 2))
}

func Test_bibao2(t *testing.T) {
	var calc Calculator
	fmt.Println(calc.Add(3, 2))
	fmt.Println(calc.Sub(3, 2))
}
