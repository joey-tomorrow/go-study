package _21010

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_waitGroup(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		wg.Done()
		wg.Add(1)
	}()
	wg.Wait()

}


var wg sync.WaitGroup

func f1() {
	time.Sleep(1 * time.Second)
	wg.Done()
}

func Test_wg(t *testing.T) {
	var i int
	for i=0;i<3;i++ {
		wg.Add(1)
		go f1()
	}
	wg.Wait()
	fmt.Println("end...")

}
