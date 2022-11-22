package gobasic

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexWoken2
	mutexWaiterShift = iota
)

func Test_goschedule(t *testing.T) {
	fmt.Println(mutexLocked)
	fmt.Println(mutexWoken)
	fmt.Println(mutexWoken2)
	fmt.Println(mutexWaiterShift)

	lock := NewSpinLock()
	for i := 0; i < 100; i++ {
		v := i
		go func(val int) {
			lock.Lock()
			fmt.Println(val)
			lock.Unlock()
		}(v)
	}

	runtime.Gosched()
}

type spinLock uint32

const maxBackoff = 64

func (sl *spinLock) Lock() {
	backoff := 1
	for !atomic.CompareAndSwapUint32((*uint32)(sl), 0, 1) {
		// Leverage the exponential backoff algorithm, see https://en.wikipedia.org/wiki/Exponential_backoff.
		for i := 0; i < backoff; i++ {
			fmt.Println("backoff waiting")
			runtime.Gosched()
		}
		if backoff < maxBackoff {
			backoff <<= 1
		}
	}
}

func (sl *spinLock) Unlock() {
	atomic.StoreUint32((*uint32)(sl), 0)
}

// NewSpinLock instantiates a spin-lock.
func NewSpinLock() sync.Locker {
	return new(spinLock)
}