package digui

import (
	"fmt"
	"testing"
)

func TestFeibonaqie(t *testing.T) {
	//fb(8)
	fmt.Println(fb(8))
}

func fb(n int) int{
	if n <= 2 {
		//fmt.Print(n)
		return n
	}

	result := fb(n-1) + fb(n - 2)
	return result
}

func dpFb(n int) int{
	if n <= 2 {
		return n
	}

	//dp := make([]int, n + 1)
	return 0
}