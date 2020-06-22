package gofunc

import (
	"fmt"
	"testing"
)

func iterator() func() int {
	a := 0
	return func() int {
		a++
		return a
	}
}

func Test_FuncGoWithState(t *testing.T) {

	iter := iterator()

	fmt.Printf("iter 1: %d\n", iter())
	fmt.Printf("iter 2: %d\n", iter())
	fmt.Printf("iter 3: %d\n", iter())
}

func Test_FuncGoWithState1(t *testing.T) {
	itera := iterator()
	iterb := iterator()

	fmt.Printf("itera 1: %d\n", itera())
	fmt.Printf("itera 2: %d\n", itera())
	fmt.Printf("itera 3: %d\n", itera())

	fmt.Printf("iterb 1: %d\n", iterb())
	fmt.Printf("iterb 2: %d\n", iterb())
	fmt.Printf("iterb 3: %d\n", iterb())
}
