package cal

import (
	"fmt"
	"testing"
)

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

func Test_reverselinklist(t *testing.T) {
	head1 := new(ListNode)
	head1.Val = 1

	head2 := new(ListNode)
	head2.Val = 2

	head3 := new(ListNode)
	head3.Val = 3

	head4 := new(ListNode)
	head4.Val = 4

	head5 := new(ListNode)
	head5.Val = 5

	head1.Next = head2
	head2.Next = head3
	head3.Next = head4
	head4.Next = head5

	group := reverse(head1)
	fmt.Println(group)
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	i := 0

	tail := head
	var pre *ListNode
	for head != nil && i < k {
		i++
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}

	if i < k {
		//再次反轉
		var preLast *ListNode
		for pre != nil {
			next := pre.Next
			pre.Next = preLast
			preLast = pre
			pre = next
		}

		return preLast
	}

	if head == nil {
		return pre
	}

	next := reverseKGroup(head, k)
	tail.Next = next

	return pre
}

func reverse(cur *ListNode) *ListNode {
	var pre *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

func reverse1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode
	cur := head
	for cur != nil {
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	return prev
}
