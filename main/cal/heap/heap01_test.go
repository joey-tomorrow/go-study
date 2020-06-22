package heap_test

import (
	"fmt"
	"go-study/main/cal/heap"
	"testing"
)

func TestHeapObject_Pop(t *testing.T) {
	hp := &heap.HeapObject{Capacity: 10, Heap: make([]int, 0)}
	for i := 0; i < 10; i++ {
		hp.Push(i)
	}

	fmt.Println(hp.Pop())
	fmt.Println(hp.Pop())
	fmt.Println(hp.Pop())
	fmt.Println(hp.Pop())
	fmt.Println(hp.Pop())
	fmt.Println(hp.Pop())
	fmt.Println(hp.Pop())
	fmt.Println(hp.Pop())
	fmt.Println(hp.Pop())
	fmt.Println(hp.Pop())
}

func Test_slice(t *testing.T) {
	slice := make([]int, 0)
	slice = append(slice, 1)
	slice = append(slice, 1)
	slice = append(slice, 1)

	fmt.Println(slice)
}
