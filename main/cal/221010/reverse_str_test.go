package _21010

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {
	src := "你好abc啊哈哈"
	dst := reverseStr([]rune(src))
	fmt.Printf("%v\n", string(dst))

	x := make([]int, 2, 10)
	fmt.Println(x)
	fmt.Println(x[2:10])
	_ = x[6:10]
	_ = x[6:]
	_ = x[2:]
}


func reverseStr(s []rune) []rune {
	for i, j := 0, len(s) - 1; i <j && j >=0; i, j = i +1 , j - 1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

func TestSliceAppend(t *testing.T) {
	errInt := new([]int)
	*errInt = append(*errInt, 1)
	fmt.Println(errInt)

	arr := make([]int, 2, 5)

	arr = append(arr, 1)
	fmt.Println("len 为", len(arr),"cap 为", cap(arr))

	arr = append(arr, 1)
	fmt.Println("len 为", len(arr),"cap 为", cap(arr))

	arr = append(arr, 1)
	fmt.Println("len 为", len(arr),"cap 为", cap(arr))

	arr = append(arr, 1)
	fmt.Println("len 为", len(arr),"cap 为", cap(arr))

}