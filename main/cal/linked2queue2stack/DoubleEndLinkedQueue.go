package linked2queue2stack

type Node struct {
	value interface{}
	Next  *Node
	Last  *Node
}

type DoubleEndLinedQueue struct {
	Head *Node
	Tail *Node
}

// 从头部插入 先进后出 可作为stack
func (d *DoubleEndLinedQueue) putFromHead(value interface{}) {
	cur := &Node{value: value}
	if d.Head == nil {
		d.Head = cur
		d.Tail = cur
	} else {
		cur.Next = d.Head
		d.Head.Last = cur
		d.Head = cur
	}
}

func (d *DoubleEndLinedQueue) popFromHead() *Node {
	if d.Head == nil {
		return nil
	}

	cur := d.Head
	if d.Head == d.Tail {
		d.Head = nil
		d.Tail = nil
	} else {
		d.Head = d.Head.Next
		cur.Next = nil
		d.Head.Last = nil
	}

	return cur
}

// 从尾部插入 先进先出 可作为queue
func (d *DoubleEndLinedQueue) putFromTail(value interface{}) {
	cur := &Node{value: value}
	if d.Head == nil {
		d.Head = cur
		d.Tail = cur
	} else {
		cur.Last = d.Tail
		d.Tail.Next = cur
		d.Tail = cur
	}
}

func (d *DoubleEndLinedQueue) popFromTail() *Node {
	if d.Head == nil {
		return nil
	}

	cur := d.Tail
	if d.Head == d.Tail {
		d.Head = nil
		d.Tail = nil
	} else {
		d.Tail = d.Tail.Last
		d.Tail.Next = nil
		cur.Last = nil
	}

	return cur
}

func (d *DoubleEndLinedQueue) isEmpty() bool {
	return d.Head == nil
}
