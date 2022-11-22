package digui

import (
	"fmt"
	"go-zero/core/mathx"
	"testing"
)

func TestWangge(t *testing.T) {
	arr := [][]int{
		{1,3,1},
		{1,5,1},
		{4,2,1},
	}

	fmt.Println(wangge3(arr))
	fmt.Println(wangge(arr))
}

/**
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
说明：每次只能向下或者向右移动一步。

举例：
输入:
arr = [
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。
 */

func wangge3(arr [][]int) int{
	m := len(arr)
	n := len(arr[0])
	if m <= 0 || n <= 0{
		return 0
	}

	dp := make([][]int, m)
	for i:= 0; i < m; i ++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = arr[0][0]

	for i := 1; i < n; i ++ {
		dp[i][0] = arr[i][0] + dp[i - 1][0]
	}

	for i := 1; i < n; i ++ {
		dp[0][i] = arr[0][i] + dp[0][i - 1]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = mathx.MinInt(dp[i-1][j], dp[i][j-1]) + arr[i][j]
		}
	}

	return dp[m - 1][n-1]
}

func wangge(arr [][]int) int{
	x := len(arr)
	y := len(arr[0])

	if x == 0 || y == 0{
		return 0
	}

	sum := 0
	if x == 1 {
		for _, element := range arr[0] {
			sum = sum + element
		}
		return sum
	}

	if y == 1 {
		for _, element := range arr {
			sum = sum + element[0]
		}
	}

	for i := 1; i < x; i ++ {
		for j := 1; j < y; j ++ {
			if arr[i][j - 1] <= arr[i - 1][j] {
				sum = sum + arr[i][j - 1]
			} else {
				sum = sum + arr[i - 1][j]
			}
		}
	}

	return sum
}


func wangge1(arr [][]int) int{
	x := len(arr)
	y := len(arr[0])

	if x == 0 || y == 0{
		return 0
	}

	dp := make([][]int, x)
	for i := range dp {
		dp[i] = make([]int, y)
	}

	dp[0][0] = arr[0][0]
	for i := 1; i < x; i++ {
		dp[i][0] = dp[i - 1][0] + arr[i][0]
	}

	for i := 1; i < y; i++ {
		dp[0][i] = dp[0][i-1] + arr[0][i]
	}

	for i := 1; i < x; i++ {
		for j := 1; j < y; j++ {
			dp[i][j] = mathx.MinInt(dp[i-1][j], dp[i][j-1]) + arr[i][j]
		}
	}
	return dp[x-1][y-1]
}
