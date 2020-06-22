package leetcode

import (
	"fmt"
	"testing"
)

/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	arr := []rune(s)
	max := 0

	//var duplicate bool
	reset := 0
	for t := 1; t < len(arr); t++ {
		tmp := 1
		for m := t - 1; m >= reset; m-- {
			if arr[t] == arr[m] {
				reset = m + 1
				goto LOOP
			} else {
				tmp++
			}
		}

		if max < tmp {
			max = tmp
		}
	LOOP:
	}

	if max == 0 {
		max = 1
	}

	return max
}

func Test_code3(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

//别人的解题
func lengthOfLongestSubstringOther(s string) int {
	if s == "" {
		return 0
	}

	max := 1
	last := -1
	for i := 0; i < len(s); i++ {
		cur := s[i]

		// 找到前一个同值的byte
		last = findLastIndexOfByte(cur, s, last, i)

		if i-last > max {
			max = i - last
		}
	}

	return max
}

// 在上一次重复下标 last, 与本次寻找的下标 cur, 之间找到是否有b的重复
// a      b    c    a     b   --> a      b    c    a     b  ...
// last  [  寻找范围 ]   cur              last               cur
func findLastIndexOfByte(b byte, s string, last, cur int) int {

	for i := cur - 1; i >= last+1; i-- {
		if s[i] == b {
			return i
		}
	}

	return last
}
