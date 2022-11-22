package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	Value interface{}
	Next  *Node
}

// 是否有環 並且找出入环的节点
func IsRingList(head *Node) *Node {
	var f, s *Node
	f = head
	s = head
	for {
		if f.Next == nil {
			fmt.Println("not a ring list")
			return nil
		}

		if f != head && f == s {
			break
		}

		f = f.Next.Next
		s = s.Next
	}

	// 快指针复位到head  快慢指针每次走一步  最后相遇的就一定是入环节点
	f = head
	for f != s {
		f = f.Next
		s = s.Next
	}

	return f
}

func generateLinkedList(l int) *Node {
	head := &Node{Value: rand.Intn(100)}
	cur := head
	for i := 0; i < l; i++ {
		cur.Next = &Node{Value: rand.Intn(100)}
		cur = cur.Next
	}

	return head
}

func IsRingListNew(head *Node) *Node {
	if head == nil {
		return nil
	}

	fast := head
	slow := head

	for {
		if fast.Next == nil {
			return nil
		}

		if fast != head && fast == slow{
			break
		}

		fast = fast.Next.Next
		slow = slow.Next
	}

	fast = head
	for fast != slow{
		fast = fast.Next
		slow = slow.Next
	}

	return fast
}

func main() {
	list := generateLinkedList(4)
	list1 := generateLinkedList(2)
	list.Next.Next.Next.Next = list1
	list1.Next.Next = list.Next.Next

	fmt.Println(IsRingListNew(list) == list.Next.Next)
	fmt.Println(IsRingList(list) == list.Next.Next)

	head := &Node{Value: 1}
	node1 := &Node{Value: 2}
	node2 := &Node{Value: 2}
	node3 := &Node{Value: 1}

	head.Next = node1
	node1.Next = node2
	node2.Next = node3

	fmt.Println(IsHuiwen(head))
	fmt.Println(IsHuiwenNew(head))
}

//是否回文
func IsHuiwen(head *Node) bool {
	var l, f *Node
	l = head
	f = head
	for {
		if f.Next == nil || f.Next.Next == nil {
			break
		}

		l = l.Next
		f = f.Next.Next
	}

	//右半部分反转
	right := l.Next
	var pre, next *Node
	for right != nil {
		next = right.Next
		right.Next = pre
		pre = right
		right = next
	}

	left := head
	has := true
	for pre != nil && left != nil {
		if pre.Value != left.Value {
			has = false
			break
		}

		pre = pre.Next
		left = left.Next
	}

	return has
}


func IsHuiwenNew(head *Node) bool {
	l := head
	f := head

	for {
		if f.Next == nil || f.Next.Next == nil {
			break
		}

		l = l.Next
		f = f.Next.Next
	}

	right := l.Next
	var pre, next *Node
	for right != nil{
		next = right.Next
		right.Next = pre
		pre = right
		right = next
	}

	left := head
	for pre != nil && left != nil{
		if pre.Value != left.Value {
			return false
		}

		left = left.Next
		pre = pre.Next
	}

	return true
}