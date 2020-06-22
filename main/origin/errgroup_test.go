package origin

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"testing"
)

func Test_ErrGroup(t *testing.T) {
	group, ctx := errgroup.WithContext(context.Background())
	for i := 0; i < 10; i++ {
		tmp := i
		group.Go(func() error {
			if tmp == 2 {
				//return fmt.Errorf("found 2!!!")
				close(ctx.Done())
			}

			fmt.Println(tmp)
			return nil
		})
	}

	//err := group.Wait()
	//assert.NotNil(t, err)
	//assert.Equal(t, err, fmt.Errorf("found 2!!!"))

	<-ctx.Done()
	fmt.Println("end")
}
