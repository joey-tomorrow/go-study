package _21010

import (
	"fmt"
	"testing"
)

/*
给定一个单链表的头节点 head,实现一个调整单链表的函数，使得每K个节点之间为一组进行逆序，并且从链表的尾部开始组起，头部剩余节点数量不够一组的不需要逆序。（不能使用队列或者栈作为辅助）
例如：
链表:1->2->3->4->5->6->7->8->null, K = 3。那么 6->7->8，3->4->5，1->2各位一组。调整后：1->2->5->4->3->8->7->6->null。其中 1，2不调整，因为不够一组。

8 7 6 5 4 3 2 1 reverse
6 7 8 3 4 5 2 1 groupForwardReverse
1 2 5 4 3 8 7 6 reverse

作者：帅地
链接：https://juejin.cn/post/6844903910386171912
来源：稀土掘金
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

type Node struct {
	value int
	Next *Node
}


func TestGroupReverseLinklist(t *testing.T) {
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

	node := groupReverse(head1, 2)
	//node := backGroupReverse(head1, 2)
	fmt.Println(node)
}

/*
===================================================正向分组逆转===================================================
*/
func otherReverse(head *Node, k int) (*Node) {
	if k == 1 {
		return reverseLinklist(head)
	}

	temp := head
	for i := 1; i < k && temp != nil; i++ {
		temp = temp.Next
	}

	//判断是否有足够的分组数量
	if temp == nil {
		return head
	}

	newTemp := temp.Next
	temp.Next = nil

	newHead := reverseLinklist(head)
	node := otherReverse(newTemp, k)
	head.Next = node
	return newHead
}

func groupReverse(head *Node, k int) *Node{
	if k == 1 {
		return reverseLinklist(head)
	}

	temp := head
	for i := 1; i < k && temp != nil; i++ {
		temp = temp.Next
	}

	//判断是否不够分组数量，如果不够直接返回不反转
	if temp == nil {
		return head
	}

	head2 := temp.Next
	temp.Next = nil

	newHead := reverseLinklist(head)
	node := groupReverse(head2, k)
	head.Next = node
	return newHead
}

func reverseLinklist(head *Node) *Node{
	if head == nil || head.Next == nil {
		return head
	}

	newNode := reverseLinklist(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newNode
}

func groupForwardReverse(head *Node, k int) *Node{
	if k == 1 {
		return reverseLinklist(head)
	}

	temp := head
	for i := 1; i < k && temp != nil; i ++ {
		temp = temp.Next
	}

	if temp == nil {
		return head
	}


	nextHead := temp.Next
	temp = nil

	newHead := reverseLinklist(head)
	nextNewHead := groupForwardReverse(nextHead, k)
	newHead.Next = nextNewHead
	return newHead
}

func reverseLinklistNew(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseLinklist(head)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

/*
===================================================正向分组逆转===================================================
*/


/*
===================================================逆向分组逆转===================================================
*/
func backGroupReverse(head *Node, k int) *Node{
	//5 > 4 > 3 > 2 > 1
	first := reverseLinklist(head)
	//4 > 5 > 2 > 3 > 1
	second := groupReverse(first, k)
	//5 > 4 > 3 > 2 > 1
	result := groupReverse(second, k)

	return result
}
