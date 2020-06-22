package origin_test

import (
	"fmt"
	"testing"
	"unsafe"
)

type Child struct {
	name string
	Mini
}

type Mini struct {
	sex int
}
type Parent struct {
	c   Child
	Age int
}

func Test_pointer(t *testing.T) {
	p := &Parent{
		c: Child{
			name: "dwliang",
		},
		Age: 22,
	}

	child := (*Child)(unsafe.Pointer(uintptr(unsafe.Pointer(p))))
	child1 := &Child{
		name: "test",
		Mini: Mini{
			sex: 1000,
		},
	}

	*child = *child1

	fmt.Println(child)
	fmt.Println(p)
}
