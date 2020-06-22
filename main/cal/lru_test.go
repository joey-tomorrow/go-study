package cal

import (
	"fmt"
	"sync"
	"testing"
)

type LRUCache struct {
	Entry map[int]*LRUNode
	Head  *LRUNode
	Tail  *LRUNode

	Capacity int
	RwLock   sync.RWMutex
}

type LRUNode struct {
	Key    int
	Value  int
	Next   *LRUNode
	Before *LRUNode
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		Entry:    make(map[int]*LRUNode),
		Capacity: capacity,
		RwLock:   sync.RWMutex{},
	}

	return cache
}

func (this *LRUCache) Get(key int) int {
	this.RwLock.Lock()
	defer this.RwLock.Unlock()
	if node, exists := this.Entry[key]; exists {
		// 是头结点 直接更新
		if this.Head == node {
			return node.Value
		}

		// 不是头结点 换到头结点
		this.swapHead(node)
		return node.Value
	}

	return -1
}

func (this *LRUCache) swapHead(node *LRUNode) {
	// 不是头结点 换到头结点
	node.Before.Next = node.Next
	if this.Tail == node {
		this.Tail = node.Before
	} else {
		node.Next.Before = node.Before
	}

	node.Next = this.Head
	this.Head.Before = node
	this.Head = node
	node.Before = nil
}

func (this *LRUCache) Put(key int, value int) {
	this.RwLock.Lock()
	defer this.RwLock.Unlock()
	// 存在
	if node, exists := this.Entry[key]; exists {
		// 是头结点 直接更新
		if this.Head == node {
			node.Value = value
			return
		}

		// 不是头结点 换到头结点
		this.swapHead(node)
		return
	}

	cur := &LRUNode{
		Key:   key,
		Value: value,
		Next:  this.Head,
	}

	if this.Head != nil {
		this.Head.Before = cur
	}

	this.Head = cur

	if this.Tail == nil {
		this.Tail = cur
	}

	this.Entry[key] = cur

	if len(this.Entry) > this.Capacity {
		oldTail := this.Tail
		tailValue := oldTail.Key

		this.Tail = oldTail.Before
		this.Tail.Next = nil
		oldTail.Before = nil

		delete(this.Entry, tailValue)
	}

	return
}

func Test_LRU(t *testing.T) {
	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // 返回  1
	cache.Put(3, 3)           // 该操作会使得密钥 2 作废
	fmt.Println(cache.Get(2)) // 返回 -1 (未找到)
	cache.Put(4, 4)           // 该操作会使得密钥 1 作废
	fmt.Println(cache.Get(1)) // 返回 -1 (未找到)
	fmt.Println(cache.Get(3)) // 返回  3
	fmt.Println(cache.Get(4)) // 返回  4
}
