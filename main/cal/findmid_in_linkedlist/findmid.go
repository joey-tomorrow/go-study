package main

import (
	"fmt"
)

type Node struct {
	Value interface{}
	Next  *Node
}

func generateLinkedList(l int) *Node {
	head := &Node{Value: 0}
	cur := head
	for i := 1; i < l; i++ {
		cur.Next = &Node{Value: i}
		cur = cur.Next
	}

	return head
}

func main() {
	list := generateLinkedList(5)
	fmt.Println(findMid(list))
}

//奇数长度 返回中间 偶数长度返回中上
func findMid(head *Node) interface{} {
	if head == nil {
		return nil
	}

	if head.Next == nil || head.Next.Next == nil {
		return head.Value
	}

	slow := head
	fast := head.Next

	for {
		if fast.Next == nil {
			break
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == nil {
			break
		}
	}

	return slow.Value

}
