package main

import (
	"fmt"
	"testing"
)

func SelectSort(arr []int, n int) []int {
	for i := 0; i < n-1; i++ {
		min := arr[i]
		idx := i
		for j := i + 1; j < n; j++ {
			if min > arr[j] {
				min = arr[j]
				idx = j
			}
		}

		if i != idx {
			swap(arr, i, idx)
		}

	}

	return arr
}

func SelectSort1(arr []int, n int) []int {
	for i := 0; i < n-1; i++ {
		minidx := i
		for j := i + 1; j < n; j++ {
			if arr[minidx] > arr[j] {
				minidx = j
			}
		}

		if i != minidx {
			swap(arr, i, minidx)
		}

	}

	return arr
}

func Test_Select(t *testing.T) {
	src := []int{6, 2, 9, 3, 5, 6}
	for _, v := range SelectSort2(src, 6) {
		fmt.Print(v)
	}
}

func SelectSort2(arr []int, n int) []int {
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i; j < n; j++ {
			if arr[minIdx] > arr[j] {
				minIdx = j
			}
		}
		if minIdx != i {
			swap(arr, i, minIdx)
		}
	}

	return arr
}
