package origin_test

import (
	"fmt"
	"sync"
	"testing"
)

func Test_WaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(tmp int) {
			defer wg.Done()
			fmt.Println("execute ", tmp)
		}(i)

	}

	wg.Wait()
}
