package _21010

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func Test_doublecheck(t *testing.T) {
	runtime.GOMAXPROCS(10)
 	once := &Once{}
 	count := 0
	for i := 0; i < 100; i++ {
		go func() {
			once.Do(func() {
				count ++
				fmt.Println(count)
			})
		}()
	}


	time.Sleep(10 * time.Second)
}

type Once struct {
	m sync.Mutex
	done uint32

}

func (o *Once) Do(f func()) {
	if o.done == 1 {
		return
	}
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		o.done = 1
		f()
	}
}