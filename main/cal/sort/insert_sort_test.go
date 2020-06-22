package main

import (
	"fmt"
	"testing"
)

func InsertSort(arr []int, n int) []int {
	for i := 1; i < n; i++ {
		for j := i; j >= 0 && j < (n-1); j-- {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			}
		}
	}

	return arr
}

func Test_Insert(t *testing.T) {
	src := []int{6, 2, 9, 3, 5, 6, 1, 3, 2, 7, 8}
	for _, v := range InsertSort1(src, 11) {
		fmt.Print(v)
	}
}

func InsertSort1(arr []int, n int) []int {
	for i := 0; i < n-1; i++ {
		for j := i; j >= 0; j-- {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			}
		}
	}

	return arr
}

//for (int i = 1; i < arr.length; i++) { // 0 ~ i 做到有序
//	for (int j = i - 1; j >= 0 && arr[j] > arr[j + 1]; j--) {
//		swap(arr, j, j + 1);
//	}
//}

// 插入排序，a表示数组，n表示数组大小
// public void insertionSort(int[] a, int n)
// {
// if (n <= 1) return;
// for (int i = 1; i < n; ++i) {
// int value = a[i];
// int j = i - 1;
// 查找插入的位置
// for (; j >= 0; --j) {
// if (a[j] > value) {
// a[j+1] = a[j]; // 数据移动
// } else { break; }
// }
// a[j+1] = value;
// 插入数据 }
// }
