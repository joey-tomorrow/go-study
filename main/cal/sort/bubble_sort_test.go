package main

import (
	"fmt"
	"testing"
)

func BubbleSort(arr []int) []int {
	for i := len(arr) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			}
		}
	}

	return arr
}

func BubbleSort1(arr []int, n int) []int {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			}
		}
	}

	return arr
}

func Test_bubbule(t *testing.T) {
	src := []int{6, 2, 9, 3, 5, 6}
	for _, v := range BubbleSort3(src, 6) {
		fmt.Print(v)
	}
}

func BubbleSort2(arr []int, n int) []int {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			}
		}
	}

	return arr
}

func BubbleSort3(arr []int, n int) []int {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			}
		}
	}

	return arr
}
