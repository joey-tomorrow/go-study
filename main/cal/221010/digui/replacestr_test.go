package digui

import (
	"fmt"
	"testing"
)


func TestReplace(t *testing.T) {
	fmt.Println(replace("dwliang", "wang"))
}

/*
给定两个单词 word1 和 word2，计算出将 word1 转换成 word2 所使用的最少操作数 。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符

示例 1:
输入: word1 = "horse", word2 = "ros"
输出: 3
解释:
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
*/

func replace(src, dest string) int{
	if dest == "" {
		return 0
	}

	dp := make([][]int, len(src))
	for i := range dp {
		dp[i] = make([]int, len(dest))
	}

	m := len(src)
	n := len(dest)

	for i := 1; i < m; i++ {
		dp[i][0] = dp[i - 1][0] + 1
	}

	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i - 1] + 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if string(src[i]) == string(dest[j]) {
				dp[i][j] = dp[i-1][j-1]
			} else {
				min := 0
				if dp[i-1][j-1] < dp[i][j-1] {
					min = dp[i-1][j-1]
				} else {
					min = dp[i][j-1]
				}

				if min > dp[i-1][j] {
					min = dp[i-1][j]
				}
				dp[i][j] = min + 1
			}
		}
	}

	return dp[m - 1][n - 1]
}

func replaceStr(src, dest string) int{
	if src == "" || dest == "" {
		return 0
	}

	m := len(src)
	n := len(dest)

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 1; i < m; i++ {
		dp[i][0] = dp[i - 1][0] + 1
	}

	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if string(src[i]) == string(dest[j]) {
				dp[i][j] = dp[i - 1][j - 1]
			} else {
				//（1）、如果把字符 word1[i] 替换成与 word2[j] 相等，则有 dp[i] [j] = dp[i-1] [j-1] + 1;

				//（2）、如果在字符串 word1末尾插入一个与 word2[j] 相等的字符，则有 dp[i] [j] = dp[i] [j-1] + 1;

				//（3）、如果把字符 word1[i] 删除，则有 dp[i] [j] = dp[i-1] [j] + 1;
			}
		}
	}

	return 0
}