package taoyi

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"testing"
	"time"
)

type A struct {
	s string
}
// 这是上面提到的 "在方法内把局部变量指针返回" 的情况

func foo(s string) *A {
	a := new(A)
	a.s = s
	return a
	//返回局部变量a,在C语言中妥妥野指针，但在go则ok，但a会逃逸到堆
}

func Test_taoyi(t *testing.T) {
	a := foo("hello")
	b := a.s + " world"
	c := b + "!"
	fmt.Println(c)
}

func Test_ErrGroup(t *testing.T) {
	group, ctx := errgroup.WithContext(context.Background())
	for i := 0; i < 10; i++ {
		tmp := i
		group.Go(func() error {
			time.Sleep(time.Duration(i) * time.Millisecond)
			if tmp == 7 {
				return fmt.Errorf("found 2!!!")
			}

			select {
			case <- ctx.Done():
				fmt.Println(fmt.Sprintf("cancel() now is : %d", tmp))
				return nil
			default:
				fmt.Println(tmp)
			}

			return nil
		})
	}

	err := group.Wait()
	fmt.Println(err)
	//assert.NotNil(t, err)
	//assert.Equal(t, err, fmt.Errorf("found 2!!!"))
}
