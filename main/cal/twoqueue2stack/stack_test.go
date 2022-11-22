package twoqueue2stack

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := NewStack()
	stack.putNew(1)
	stack.putNew(2)
	stack.putNew(3)
	stack.putNew(4)

	fmt.Println(stack.popNew())

	stack.putNew(5)
	stack.putNew(6)

	fmt.Println(stack.popNew())
	fmt.Println(stack.popNew())
	fmt.Println(stack.popNew())

	stack.putNew(8)
	fmt.Println(stack.popNew())
	fmt.Println(stack.popNew())
	fmt.Println(stack.popNew())

	strings.Contains("abcdefghijklmabcdefghijklmnabcdefghijklmnabcdefghijklmnabcdefghijklmnabcdefghijklmnabcdefghijklmnabcdefghijklmnndwlaing000000000000123214555", "123")

	//name := "dwliang"
	s := new(sku)
	//s.Name = &name

	if s.Name == nil || *s.Name != "" {
		fmt.Println(s.Name)
	}
}

type sku struct {
	Name *string
}
