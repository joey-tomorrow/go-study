package gofunc

import (
	"fmt"
	"testing"
)

//https://studygolang.com/articles/12703?fr=sidebar/*&^%$

type stateChanger func() int

func new() (stateChanger, stateChanger) {
	a := 0
	return func() int {
			a++
			return a
		},
		func() int {
			a--
			return a
		}
}

func Test_bibao(t *testing.T) {
	inc, dec := new()

	fmt.Printf("inc 1: %d\n", inc())
	fmt.Printf("inc 2: %d\n", inc())
	fmt.Printf("inc 3: %d\n", inc())

	fmt.Printf("dec 1: %d\n", dec())
	fmt.Printf("dec 2: %d\n", dec())
	fmt.Printf("dec 3: %d\n", dec())
}
