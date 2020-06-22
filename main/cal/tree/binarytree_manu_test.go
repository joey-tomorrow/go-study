package tree_test

import (
	"fmt"
	"go-study/main/cal/linked2queue2stack"
	"testing"
)

type Node struct {
	Value interface{}
	Left  *Node
	Right *Node
}

func preManu(node *Node) {
	if node == nil {
		return
	}

	stack := linked2queue2stack.NewStack()
	stack.Put(node)

	for !stack.IsEmpty() {
		pop := stack.Pop().(*Node)
		fmt.Println(pop.Value)

		if pop.Right != nil {
			stack.Put(pop.Right)
		}

		if pop.Left != nil {
			stack.Put(pop.Left)
		}
	}

}

func pre(node *Node) {
	if node == nil {
		return
	}

	fmt.Println(node.Value)
	pre(node.Left)
	pre(node.Right)
}

func mid(node *Node) {
	if node == nil {
		return
	}

	mid(node.Left)
	fmt.Println(node.Value)
	mid(node.Right)
}

func post(node *Node) {
	if node == nil {
		return
	}

	post(node.Left)
	post(node.Right)
	fmt.Println(node.Value)
}

func Test_premanu(t *testing.T) {
	node1 := &Node{Value: 1}
	node2 := &Node{Value: 2}
	node3 := &Node{Value: 3}
	node4 := &Node{Value: 4}
	node5 := &Node{Value: 5}
	node6 := &Node{Value: 6}
	node7 := &Node{Value: 7}

	node1.Left = node2
	node1.Right = node3

	node2.Left = node4
	node2.Right = node5

	node3.Left = node6
	node3.Right = node7

	preManu(node1)
	fmt.Println("------------------------------------")
	pre(node1)
	fmt.Println("------------------------------------")
}
