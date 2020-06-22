package cal

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Stack struct {
	top    *Node
	length int
}

type Node struct {
	val  interface{}
	prev *Node
}

// 温度
type Temperatures struct {
	Temp int
	Day  int
}

func New() *Stack {
	return &Stack{nil, 0}
}

// 入栈
func (s *Stack) Push(value interface{}) {
	n := &Node{value, s.top}
	s.top = n
	s.length++
}

// 查看栈顶元素
func (s *Stack) Peek() interface{} {
	if s.length == 0 {
		return ""
	}
	return s.top.val
}

// 出栈
func (s *Stack) Pop() interface{} {
	if s.length == 0 {
		return ""
	}
	n := s.top
	s.top = n.prev
	s.length--
	return n.val
}

func isValid(s string) bool {
	right := make(map[string]string)
	right[")"] = "("
	right["]"] = "["
	right["}"] = "{"
	stack := New()
	for _, v := range s {
		if l, ok := right[string(v)]; ok {
			top := stack.Pop()
			if l != top {
				return false
			}
		} else {
			stack.Push(string(v))
		}
	}

	if stack.length > 0 {
		return false
	}
	return true
}

// 左右括号匹配
func Test_rightleft(t *testing.T) {
	s := "["
	assert.Equal(t, isValid(s), false)
}

// 测试温度差  leetcode-739
func Test_wenducha(t *testing.T) {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}

/*
根据每日 气温 列表，请重新生成一个列表，对应位置的输入是你需要再等待多久温度才会升高超过该日的天数。如果之后都不会升高，请在该位置用 0 来代替。

例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]。

提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数。

*/
func dailyTemperatures(T []int) []int {
	stack := New()
	result := make([]int, len(T))
	for idx, v := range T {

		for {
			if stack.top == nil {
				stack.Push(idx)
				break
			}
			topIdx := stack.top.val.(int)
			if v > T[topIdx] {
				result[topIdx] = idx - topIdx
				stack.Pop()
			} else {
				stack.Push(idx)
				break
			}
		}

	}

	return result
}
