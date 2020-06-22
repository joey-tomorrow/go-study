package leetcode

import (
	"fmt"
	"testing"
)

/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := new(ListNode)
	h3 := l3
	add := 0
	for l1 != nil || l2 != nil || add > 0 {
		h3.Next = new(ListNode)
		h3 = h3.Next

		tmp := add
		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}

		add = tmp / 10

		h3.Val = tmp % 10

	}

	return l3.Next
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	node1 := l1
	node2 := l2
	l3 := new(ListNode)
	node3 := l3
	a := 0
	for node1 != nil || node2 != nil || a > 0 {
		node3.Next = new(ListNode)
		node3 = node3.Next
		b := 0
		c := 0
		if node1 != nil {
			b = node1.Val
		}
		if node2 != nil {
			c = node2.Val
		}
		node3.Val = (a + b + c) % 10
		a = (a + b + c) / 10
		if node1 != nil {
			node1 = node1.Next
		}
		if node2 != nil {
			node2 = node2.Next
		}
	}
	return l3.Next
}

func Test_code2(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l11 := &ListNode{Val: 7}
	l12 := &ListNode{Val: 3}

	l1.Next = l11
	l11.Next = l12

	l2 := &ListNode{Val: 3}
	l21 := &ListNode{Val: 4}
	l22 := &ListNode{Val: 5}

	l2.Next = l21
	l21.Next = l22

	//numbers1 := addTwoNumbers1(l1, l2)
	//fmt.Println(numbers1)

	numbers2 := addTwoNumbers(l1, l2)
	fmt.Println(numbers2)
}
