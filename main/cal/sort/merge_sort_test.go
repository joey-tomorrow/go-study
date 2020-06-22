package main_test

import (
	"fmt"
	"testing"
)

//go test -v -test.run TestMergerSort
func TestMergerSort(t *testing.T) {
	fmt.Println(len(""))
	if true {
		fmt.Println("自己写的代码")
		// 测试代码
		arr := []int{9, 8, 7, 6, 5, 1, 2, 3, 4, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
		fmt.Println(arr)
		mergerSort_MyCode(arr, 0, len(arr))
		fmt.Println(arr)
	}
	fmt.Println("-----------------")
	if true {
		fmt.Println("网上复制的代码")
		// 测试代码
		arr := []int{9, 8, 7, 6, 5, 1, 2, 3, 4, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
		fmt.Println(arr)
		mergerSort_CopyCode(arr, 0, len(arr))
		fmt.Println(arr)
	}
}

// mergerSort_MyCode自己写的代码 步骤：1.判断长度是否大于1；2.取中值；3.递归排序两个；4.循环左数组，如果逆序，交换并冒泡排序右数组一次。
func mergerSort_MyCode(arr []int, left, right int) {
	if right-left <= 1 { //长度小于等于1，不需要排序，直接返回
		return
	}

	mid := left + (right-left)>>1      //取中值不会溢出
	mergerSort_MyCode(arr, left, mid)  //左排序
	mergerSort_MyCode(arr, mid, right) //右排序

	for i := left; i < mid; i++ { //循环左边
		if arr[i] > arr[mid] { //如果倒序，需要交换
			arr[i], arr[mid] = arr[mid], arr[i]
			//mid到right冒泡排序一轮，如果有序，中断循环
			for m := mid; m+1 < right; m++ {
				if arr[m] > arr[m+1] {
					arr[m], arr[m+1] = arr[m+1], arr[m]
				} else {
					break
				}
			}
		}
	}
}

// mergerSort_CopyCode网上复制的代码
func mergerSort_CopyCode(arr []int, left, right int) {
	if right-left <= 1 {
		return
	}

	//c := (a + b) / 2 //取中值可能溢出
	mid := left + (right-left)>>1        //取中值不会溢出
	mergerSort_CopyCode(arr, left, mid)  //左排序
	mergerSort_CopyCode(arr, mid, right) //右排序

	arrLeft := make([]int, mid-left)
	arrRight := make([]int, right-mid)
	copy(arrLeft, arr[left:mid])   //复制数组左边
	copy(arrRight, arr[mid:right]) //复制数组右边
	i := 0
	j := 0
	for k := left; k < right; k++ {
		if i >= mid-left { //无左边
			arr[k] = arrRight[j]
			j++
		} else if j >= right-mid { //无右边
			arr[k] = arrLeft[i]
			i++
		} else if arrLeft[i] <= arrRight[j] { //有序
			arr[k] = arrLeft[i]
			i++
		} else {
			arr[k] = arrRight[j] //倒序
			j++
		}
	}
}
