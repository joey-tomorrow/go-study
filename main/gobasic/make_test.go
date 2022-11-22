package gobasic

import (
	"context"
	"fmt"
	"testing"
)

type parent struct {
	num uint64
}

func testCtx(ctx *context.Context) {
	fmt.Println(ctx)
}

func TestMakeNew(t *testing.T) {


}

type Stream struct {
	data []int
}

func (s Stream) Map(f func(int) int) Stream {
	res := make([]int, 0, len(s.data))
	for _, item := range s.data {
		res = append(res, f(item))
	}
	return Stream{res}
}

func (s Stream) Filter(f func(int) bool) Stream {
	res := []int{}
	for _, item := range s.data {
		if f(item) {
			res = append(res, item)
		}
	}
	return Stream{res}
}

func TestMapReduce(t *testing.T) {

	stream := Stream{[]int{1, 3, 5, 7, 8}}
	fmt.Println(stream.data)
	stream = stream.
		Map(func(i int) int { return i + 1 }).
		Filter(func(i int) bool { return i%2 == 0 }).
		Filter(func(i int) bool { return i >5})
	fmt.Println(stream.data)
}