package everyday

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

func Test_cancel(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go func() {
		wg.Add(2)
		go watch(ctx, 1)
		go watch(ctx, 2)
		wg.Wait()
	}()

	select {
	case <-ctx.Done():
		fmt.Printf("watch %d %s\n", 0, ctx.Err())
	}

	fmt.Println("finished")
}

func watch(ctx context.Context, flag int) {
	defer wg.Done()

	func() {
		fmt.Printf("doing something flag:%d\n", flag)
		time.Sleep(50 * time.Second)
		fmt.Println("finished flag:", flag)
	}()
}
