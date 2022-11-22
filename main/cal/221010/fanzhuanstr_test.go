package _21010

import (
	"fmt"
	"testing"
)

/**
问题描述
请实现⼀个算法，在不使⽤【额外数据结构和储存空间】的情况下，翻转⼀个给定的字
符串(可以使⽤单个过程变量)。
给定⼀个string，请返回⼀个string，为翻转后的字符串。保证字符串的⻓度⼩于等于
5000。
 */
func Test_fanzhuanstr(t *testing.T) {
	str := "abcdefghijklmno"
	fmt.Println(reverse(str))
}

func reverse(s string) string{
	l := len(s)
	if len(s) > 5000 {
		return ""
	}

	str := []rune(s)

	fmt.Println(l/2)
	for i := 0; i < l/2; i ++ {
		str[i],str[l - i - 1] = str[l - i - 1], str[i]
	}

	return string(str)
}


type Param map[string]interface{}
type show struct {
	Param
}


type student struct {
	Name string
}

func zhoujielun(v interface{}) {
	switch msg := v.(type) {
	case *student, student:
		fmt.Println(msg.(student).Name)
	}
}

func Test_code1(t *testing.T) {
	zhoujielun(&student{Name: "dwliang"})
	s := new(show)
	s.Param["1"] = 100
	fmt.Println(s.Param["1"])
}


type People struct {
	Name string
}

func (p *People) String1() string {
	return fmt.Sprintf("print: %v", p)
}

func Test_string(t *testing.T) {
	p := &People{}
	p.String1()
}