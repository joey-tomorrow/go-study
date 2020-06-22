package main

import (
	"fmt"
	"testing"
)

func partion(arr []int, L, R int) int {
	if L == R {
		return L
	}

	if L > R {
		return -1
	}

	less := L - 1
	index := L

	for index < R {
		if arr[index] < arr[R] {
			less++
			swap(arr, index, less)
		}
		index++
	}

	less++
	swap(arr, less, R)

	return less
}

func partionSort(arr []int, L, R int) {
	if L >= R {
		return
	}

	mid := partion1(arr, L, R)
	partionSort(arr, L, mid-1)
	partionSort(arr, mid+1, R)
}

//
//func swap(arr []int, i, j int) {
//	arr[i], arr[j] = arr[j], arr[i]
//}

func Test_partion(t *testing.T) {
	src := []int{6, 2, 9, 3, 5, 6, 1, 4}
	partionSort(src, 0, 7)
	for _, v := range src {
		fmt.Print(v)
	}
}

func partion1(arr []int, L, R int) int {
	if L == R {
		return L
	}

	if L > R {
		return -1
	}

	idx := L
	less := L - 1

	for idx < R {
		if arr[idx] < arr[R] {
			less++
			swap(arr, less, idx)
		}
		idx++
	}

	less++
	swap(arr, less, R)
	return less
}
