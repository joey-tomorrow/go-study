package main

import (
	"fmt"
	"math"
)

type Node struct {
	Value interface{}
	Next  *Node
}

// 找出两个链表的重叠点
func findIntersect(node1, node2 *Node) *Node {
	count := 0
	head1 := node1
	head2 := node2
	for head1 != nil {
		count++
		head1 = head1.Next
	}

	for head2 != nil {
		count--
		head2 = head2.Next
	}

	var (
		lNode, sNode *Node
	)

	//誰長，誰是lNode
	if count > 0 {
		lNode = node1
		sNode = node2
	} else {
		sNode = node1
		lNode = node2
	}

	count1 := int(math.Abs(float64(count)))
	for count1 > 0 {
		lNode = lNode.Next
		count1--
	}

	for lNode != sNode {
		lNode = lNode.Next
		sNode = sNode.Next
	}

	return lNode
}

func generateLinkedList(l int) *Node {
	head := &Node{Value: 3}
	cur := head
	for i := 0; i < l; i++ {
		cur.Next = &Node{Value: i}
		cur = cur.Next
	}

	return head
}

func main() {
	node1 := generateLinkedList(5)
	node2 := generateLinkedList(3)
	node2.Next.Next.Next = node1

	fmt.Println("expect:", node1.Value)

	fmt.Println(findIntersect(node1, node2).Value)
}
