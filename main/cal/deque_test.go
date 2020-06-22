package cal_test

import (
	"fmt"
	"testing"
)

// 链表的一个节点
type ListNode struct {
	prev  *ListNode   // 前一个节点
	next  *ListNode   // 后一个节点
	value interface{} // 数据
}

// 创建一个节点
func NewListNode(value interface{}) (listNode *ListNode) {
	listNode = &ListNode{
		value: value,
	}

	return
}

// 当前节点的前一个节点
func (n *ListNode) Prev() (prev *ListNode) {
	prev = n.prev

	return
}

// 当前节点的前一个节点
func (n *ListNode) Next() (next *ListNode) {
	next = n.next

	return
}

// 获取节点的值
func (n *ListNode) GetValue() (value interface{}) {
	if n == nil {

		return
	}
	value = n.value

	return
}

// 链表
type List struct {
	head *ListNode // 表头节点
	tail *ListNode // 表尾节点
	len  int       // 链表的长度
}

// 创建一个空链表
func NewList() (list *List) {
	list = &List{}
	return
}

// 返回链表头节点
func (l *List) Head() (head *ListNode) {
	head = l.head

	return
}

// 返回链表尾节点
func (l *List) Tail() (tail *ListNode) {
	tail = l.tail

	return
}

// 返回链表长度
func (l *List) Len() (len int) {
	len = l.len

	return
}

// 在链表的右边插入一个元素
func (l *List) RPush(value interface{}) {

	node := NewListNode(value)

	// 链表未空的时候
	if l.Len() == 0 {
		l.head = node
		l.tail = node
	} else {
		tail := l.tail
		tail.next = node
		node.prev = tail

		l.tail = node
	}

	l.len = l.len + 1

	return
}

// 从链表左边取出一个节点
func (l *List) LPop() (node *ListNode) {

	// 数据为空
	if l.len == 0 {

		return
	}

	node = l.head

	if node.next == nil {
		// 链表未空
		l.head = nil
		l.tail = nil
	} else {
		l.head = node.next
	}
	l.len = l.len - 1

	return
}

// 从链表右边取出一个节点
func (l *List) RPop() (node *ListNode) {

	// 数据为空
	if l.len == 0 {

		return
	}

	node = l.tail
	l.tail = node.prev
	//if node.next == nil {
	//	// 链表未空
	//	l.head = nil
	//	l.tail = nil
	//} else {
	//
	//}
	l.len = l.len - 1

	return
}

func Test_deque(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println(maxSlidingWindow(nums, 3))
}

func maxSlidingWindow(nums []int, k int) []int {
	d := NewList()
	windowCnt := len(nums) - k + 1
	result := make([]int, 0, windowCnt)
	for i := 0; i < windowCnt; i++ {

		for j := 0; j < k; j++ {
			if i == 0 && j == 0 {
				d.RPush(nums[i])
				continue
			}

			max := nums[i+j]
			for t := i + j - 1; t >= i; t-- {
				if max > nums[t] {
					d.RPop()
					if t == i {
						d.RPush(i + j)
					}
				} else {
					d.RPush(t + 1)
					break
				}
			}
		}

		if d.Head().value.(int) < i {
			d.LPop()
			result = append(result, nums[d.LPop().value.(int)])
		} else {
			result = append(result, nums[d.Head().value.(int)])
		}
	}

	return result
}

//func maxSlidingWindow1(nums []int, k int) []int {
//	d:= NewList()
//	windowCnt := len(nums) - k + 1
//	result := make([]int, 0, windowCnt)
//	for i := k; i < len(nums); i ++ {
//
//	}
//
//
//	return result
//}
