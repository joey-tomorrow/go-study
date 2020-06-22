package main

import (
	"fmt"
	"testing"
)

// 荷兰国旗问题 以r为基准  小于r的在最左边 相等的在中间 大于的在右边
func netherlandsFlagSort(arr []int, l, r int) (int, int) {
	if l == r {
		return l, r
	}

	idx := l
	less := l - 1
	more := r
	for {
		if idx == more {
			break
		}

		if arr[idx] == arr[r] {
			idx++
		} else if arr[idx] < arr[r] {
			less++
			swap(arr, idx, less)
			idx++
		} else {
			more--
			swap(arr, idx, more)
		}
	}

	swap(arr, more, r)
	for _, v := range arr {
		fmt.Println(v)
	}

	return less + 1, more

}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func Test_sortHerland(t *testing.T) {
	src := []int{6, 2, 9, 3, 5, 6}
	fmt.Println(netherlandsFlagSort1(src, 0, 5))
}

// 荷兰国旗问题 以r为基准  小于r的在最左边 相等的在中间 大于的在右边
func netherlandsFlagSort1(arr []int, l, r int) (int, int) {
	if l == r {
		return l, r
	}

	if l > r {
		return 0, 0
	}

	less := l - 1
	index := l
	more := r
	for index < more {
		if arr[index] < arr[r] {
			less++
			swap(arr, index, less)
			index++
		} else if arr[index] == arr[r] {
			index++
		} else {
			more--
			swap(arr, more, index)
		}
	}

	swap(arr, index, r)
	for _, v := range arr {
		fmt.Println(v)
	}

	less++
	return less, more
}
