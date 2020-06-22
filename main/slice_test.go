package main_test

import (
	"fmt"
	"testing"
)

func Test_Slice(t *testing.T) {
	newSlice := new([]int)
	fmt.Println(len(*newSlice))
	*newSlice = append(*newSlice, 1)
	for k, v := range *newSlice {
		fmt.Printf("K:%d V:%d", k, v)
	}

	makeSlice := make([]int, 0)
	makeSlice[0] = 1
}
