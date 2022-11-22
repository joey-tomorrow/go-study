package digui

import (
	"fmt"
	"testing"
)

func TestQingwa(t *testing.T) {
	fmt.Println(skip(8))
	fmt.Println(f(8))
	fmt.Println(dpSkip(8))
}

func skip(n int) int{
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	return skip(n-1) + skip(n-2)
}

func f(n int) int{
	if n <= 2 {
		return n
	}

	f1 := 1
	f2 := 2
	sum := 0

	for i := 3; i <= n; i++ {
		sum = f1 + f2
		f1 = f2
		f2 = sum
	}

	return sum
}

//动态规划
func dpSkip(n int) int{
	if n <= 2 {
		return n
	}

	// 先创建一个数组来保存历史数据
	dp := make([]int, n+1)

	// 给出初始值
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2

	// 通过关系式来计算出 dp[n]
	for i := 3; i <= n; i++ {
		dp[i] = dp[i - 1] + dp[i - 2]
	}

	return dp[n]
}
