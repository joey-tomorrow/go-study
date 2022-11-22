package main

import (
	"fmt"
	"go-study/main/cal/linked2queue2stack"
)

func main() {
	printStack()
	fmt.Println("===============")
	printQueue()
	fmt.Println("===============")
	printQueue1()
}

func printStack() {
	stack := &linked2queue2stack.Stack{
		&linked2queue2stack.DoubleEndLinedQueue{},
	}

	for i := 0; i < 10; i++ {
		stack.Put(i)
	}

	for i := 0; i < 10; i++ {
		fmt.Print(stack.Pop())
	}
}

func printQueue() {
	queue := &linked2queue2stack.Queue{
		&linked2queue2stack.DoubleEndLinedQueue{},
	}

	for i := 0; i < 10; i++ {
		queue.Put(i)
	}

	for i := 0; i < 10; i++ {
		fmt.Print(queue.Pop())
	}
}

func printQueue1() {
	queue := &linked2queue2stack.Queue1{
		&linked2queue2stack.DoubleEndLinedQueue{},
	}

	for i := 0; i < 10; i++ {
		queue.Put(i)
	}

	for i := 0; i < 10; i++ {
		fmt.Print(queue.Pop())
	}
}
