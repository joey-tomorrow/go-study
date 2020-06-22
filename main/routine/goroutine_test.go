package routine

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func Benchmark_Goroutine(t *testing.B) {

	runtime.GOMAXPROCS(1)

	currentTime := time.Now()
	fmt.Println(currentTime)

	go func() {
		for i := 0; i < 1000000; i++ {
			runtime.Gosched()
		}
	}()

	for i := 0; i < 1000000; i++ {
		runtime.Gosched()
	}

	currentTime = time.Now()
	fmt.Println(currentTime)
}
