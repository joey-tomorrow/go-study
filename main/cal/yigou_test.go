package cal

import (
	"fmt"
	"testing"
)

/**
给定两个字符串 s 和 t，编写一个函数来判断 t 是否是 s 的字母异位词。



说明：你可以假设字符串只包含小写字母。



示例 1

输入: s = "anagram", t = "nagaram"

输出: true



示例 2

输入: s = "rat", t = "car"

输出: false
*/
func Test_yigou(t *testing.T) {
	fmt.Println(answer1("dwliang", "dlwiang"))
	fmt.Println(answer2("dwliang", "dlwiang"))
}

// 解法1：使用两个数组存放26个英文字母
func answer1(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	srcArr := make([]rune, 26)
	destArr := make([]rune, 26)
	for _, v := range s {
		index := int(v) - 97
		current := srcArr[index]
		srcArr[index] = current + 1
	}

	for _, v := range t {
		index := int(v) - 97
		current := destArr[index]
		destArr[index] = current + 1
	}

	for i := 0; i < 26; i++ {
		if srcArr[i] != destArr[i] {
			return false
		}
	}

	return true
}

// 解法2：使用一个数组
func answer2(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	arr := make([]rune, 26)
	for _, v := range s {
		index := int(v) - 97
		current := arr[index]
		arr[index] = current + 1
	}

	for _, v := range t {
		index := int(v) - 97
		current := arr[index]
		arr[index] = current - 1
	}

	for _, v := range arr {
		if v > 0 {
			return false
		}
	}

	return true
}
