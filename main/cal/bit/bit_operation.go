package main

import "fmt"

//交换两个数
func swapWithoutTmp(a, b int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println(a, " ", b)
}

//数组中只有一个数个数为奇数  找出这个数
func getOdd(arr []int) int {
	var eor int
	for _, v := range arr {
		eor = eor ^ v
	}

	return eor
}

//数组中有2个数个数为奇数  找出这两个数
func get2Odd(arr []int) (int, int) {
	var eor int
	for _, v := range arr {
		eor = eor ^ v
	}

	// eor = a ^ b
	// eor != 0
	// eor必然有一个位置上是1
	// 0110010000
	// 0000010000
	right := eor & (^eor + 1)

	// 因为是异或 a b 必然有一个在right位置上等于0 另一个等于1
	// 重新开始迭代 找出与right按位与等于0的数（这里就排除了该位置为1的a或b） 然后全集异或 必然会找出a或b
	var eor1 int
	for _, v := range arr {
		if v&right == 0 {
			eor1 ^= v
		}
	}

	return eor1, eor1 ^ eor
}

//获取第一个不为1的位置
func getLasted1(a int) int {
	return 0
}

func main() {
	swapWithoutTmp(1, 1)
	fmt.Println(getOdd([]int{3, 4, 2, 6, 10, 12, 4, 2, 6, 10, 12, 4, 2, 6, 10, 12, 4, 2, 6, 10, 12, 3, 3}) == 3)
	fmt.Println(get2Odd([]int{3, 4, 2, 6, 10, 12, 4, 2, 6, 10, 12, 4, 2, 6, 10, 12, 4, 2, 6, 10, 12, 3, 3, 1, 1, 1, 1, 1}))
	fmt.Println(get2OddNew([]int{3, 4, 2, 6, 10, 12, 4, 2, 6, 10, 12, 4, 2, 6, 10, 12, 4, 2, 6, 10, 12, 3, 3, 1, 1, 1, 1, 1}))
}

//数组中有2个数个数为奇数  找出这两个数
func get2OddNew(arr []int) (int, int) {
	eor := 0

	for _, v := range arr {
		eor = eor ^ v
	}

	rightOne := eor & (^eor + 1)

	eor1 := 0

	for _, v := range arr {
		if v&rightOne == 0 {
			eor1 = eor1 ^ v
		}
	}

	return eor1, eor ^ eor1
}
