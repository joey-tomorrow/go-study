package _21010

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func Test_channel(t *testing.T) {
	fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	var ch chan int
	go func() {
		//ch = make(chan int, 1)
		ch <- 1
		fmt.Println("111")
	}()
	go func(ch chan int) {
		//time.Sleep(time.Second)
		<-ch
		fmt.Println("222")

	}(ch)

	ch <- 1
	c := time.Tick(1 * time.Second)
	for range c {
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	}

}

func TestChannel2(t *testing.T) {
	var ch chan int
	var count int
	go func() {
		ch <- 1
	}()
	go func() {
		count++
		//close(ch)
	}()
	fmt.Println(count)
	<-ch
	fmt.Println(count)

}