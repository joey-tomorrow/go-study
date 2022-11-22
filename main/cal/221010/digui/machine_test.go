package digui

import (
	"fmt"
	"testing"
)

/**
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

问总共有多少条不同的路径？
**/

func TestMachine(t *testing.T) {
	fmt.Println(dp(3, 7))
}

func dp(x int, y int) int {
	if x <= 0 || y <= 0 {
		return 0
	}

	dp := make([][]int, x)

	for i := range dp {
		dp[i] = make([]int, y)
	}

	for i := 0; i < x; i++ {
		dp[i][0] = 1
	}

	for i := 0; i < y; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < x; i++ {
		for j := 1; j < y; j++ {
			dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
		}
	}

	return dp[x-1][y-1]
}

