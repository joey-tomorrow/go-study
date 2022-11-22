package taoyi

import (
	"fmt"
	"testing"
)

type B struct {
	s string
}

type C struct {
	s string
}

func escape1(c *C) B {
	return escape(c)
}

func escape(c *C) B {
	b := new(B)
	b.s = c.s
	return *b
	//返回局部变量a,在C语言中妥妥野指针，但在go则ok，但a会逃逸到堆
}

func sum(arr []int) int {
	s := 0
	for i := 0; i < len(arr); i++ {
		s += arr[i]
	}
	return s
}

func sumTest(slic []*B) {
	l := len(slic)
	fmt.Println(l)

	if l > 1 {
		slic = append(slic, new(B))
	}
}

//go build -gcflags=-m main.go
func Test_taoyi2(t *testing.T) {
	tmp := []*B{
		&B{
			"s1",
		},
		&B{
			"s2",
		},
	}

	strings := make([]string, 2)
	strings = append(strings, "1")
	strings[0] = "str"
	//s1 := "1"
	//s2 := "2"
	//strings = append(strings, s1, s2)
	str := strings[0]

	sumTest(tmp)
	fmt.Println(str)
}

//func Test_ErrGroup(t *testing.T) {
//	group, ctx := errgroup.WithContext(context.Background())
//	for i := 0; i < 10; i++ {
//		tmp := i
//		group.Go(func() error {
//			time.Sleep(time.Duration(i) * time.Millisecond)
//			if tmp == 7 {
//				return fmt.Errorf("found 2!!!")
//			}
//
//			select {
//			case <- ctx.Done():
//				fmt.Println(fmt.Sprintf("cancel() now is : %d", tmp))
//				return nil
//			default:
//				fmt.Println(tmp)
//			}
//
//			return nil
//		})
//	}
//
//	err := group.Wait()
//	fmt.Println(err)
//	//assert.NotNil(t, err)
//	//assert.Equal(t, err, fmt.Errorf("found 2!!!"))
//}
