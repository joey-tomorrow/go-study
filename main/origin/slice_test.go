package origin_test

import (
	"fmt"
	"go-study/main/origin"
	"testing"
)

func Benchmark1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{1, 43, 61, 4, 5, 2, 8, 54, 21, 123, 123}
		origin.SliceMine(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
	}
}

func Benchmark(b *testing.B) {
	arr := []int{1, 43, 61, 4, 5, 2, 8, 54, 21, 123, 123}
	origin.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	fmt.Println(arr)

	for i := 0; i < b.N; i++ {
		arr := []int{1, 43, 61, 4, 5, 2, 8, 54, 21, 123, 123}
		origin.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})

	}
}

func Benchmark2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []string{"1", "43", "61", "4", "5", "2", "8", "54", "21", "123", "123"}
		origin.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
	}
}
