package digui

import (
	"fmt"
	"testing"
)

type Node struct {
	value int
	Next *Node
}

func TestReverse(t *testing.T) {
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

	result := reverse(head1)
	fmt.Print(result)
}

func reverse(head *Node) *Node{
	if head == nil || head.Next == nil {
		return head
	}

	newNode := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newNode
}

func reverse1(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	newNode := reverse1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newNode
}

func reverse2(head *Node) *Node {
	if head == nil || head.Next == nil{
		return head
	}

	nodeNew := reverse2(head)
	head.Next.Next = head
	head.Next = nil
	return nodeNew
}
