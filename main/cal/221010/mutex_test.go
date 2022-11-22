package _21010

import (
	"fmt"
	"sync"
	"testing"
)

var mu sync.Mutex
var chain string

func Test_mutex(t *testing.T) {
	chain = "main"
	A()
	fmt.Println(chain)
}

func A() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + " --> A"
	B()
}

func B() {
	chain = chain + " --> B"
	C()
}

func C() {
	//mu.Lock()
	//defer mu.Unlock()
	chain = chain + " --> C"
}


type MyMutex struct {
	count int
	sync.Mutex
}

func Test_myMutex(t *testing.T) {
	var mu MyMutex
	mu.Lock()
	var mu2 = mu
	mu.count++
	mu.Unlock()
	mu2.Lock()
	mu2.count++
	mu2.Unlock()
	fmt.Println(mu.count, mu2.count)
}