package _21010

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_rwmutex(t *testing.T) {
	go A1()
	time.Sleep(2 * time.Second)
	rwmu.Lock()
	defer rwmu.Unlock()
	count++
	fmt.Println(count)
}

var rwmu sync.RWMutex
var count int

func A1() {
	rwmu.RLock()
	defer rwmu.RUnlock()
	B1()
}

func B1() {
	time.Sleep(5 * time.Second)
	fmt.Println("b1()")
	C1()
}

func C1() {
	rwmu.RLock()
	defer rwmu.RUnlock()
	fmt.Println("c1()")
}