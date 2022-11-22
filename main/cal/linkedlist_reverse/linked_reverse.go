package main

import "fmt"

type Node struct {
	value int
	Next  *Node
}

func main() {
	head1 := new(Node)
	head1.value = 1

	head2 := new(Node)
	head2.value = 2

	head3 := new(Node)
	head3.value = 3

	head4 := new(Node)
	head4.value = 4

	head5 := new(Node)
	head5.value = 5

	head1.Next = head2
	head2.Next = head3
	head3.Next = head4
	head4.Next = head5

	result := reverse5(head1)
	fmt.Print(result)

}

func reverse5(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	newNode := reverse5(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newNode
}

func reverse4(head *Node) *Node {
	var pre *Node
	var next *Node

	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}

	return pre
}

func reverse(head *Node) *Node {
	var pre *Node
	var next *Node

	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}

	return pre
}

func reverse1(head *Node) *Node {
	var pre *Node
	var next *Node

	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

func reverse2(head *Node) *Node {
	var pre, next *Node

	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}

	return pre
}

func reverse3(head *Node) *Node {
	var pre *Node
	var next *Node

	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}

	return pre
}
