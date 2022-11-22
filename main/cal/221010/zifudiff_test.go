package _21010

import (
	"fmt"
	"strings"
	"testing"
)

func Test_zifudiff(t *testing.T) {
	str := "abcdefgabc"
	fmt.Println(isStrAllUnique(str))

	for idx, v := range str {
		fmt.Println( idx)
		fmt.Println(strings.LastIndex(str, string(v)))
	}

	str = "abcdefg"
	fmt.Println(isStrAllUnique(str))


}

func isStrAllUnique(s string) bool{
	if strings.Count(s, "") > 3000 {
		return false
	}

	for _, item := range s {
		if item > 127 {
			return false
		}

		if strings.Count(s, string(item)) > 1 {
			return false
		}
	}

	return true
}

