package taoyi

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

/*
那么如果想要在底层转换⼆者，只需要把 StringHeader 的地址强转成 SliceHeader 就
⾏。那么go有个很强的包叫 unsafe 。
1. unsafe.Pointer(&a) ⽅法可以得到变量a的地址。
2. (*reflect.StringHeader)(unsafe.Pointer(&a)) 可以把字符串a转成底层结构的形式。
3. (*[]byte)(unsafe.Pointer(&ssh)) 可以把ssh底层结构体转成byte的切⽚的指针。
4. 再通过  * 转为指针指向的实际内容。
*/
func Test_memory_copy(t *testing.T) {
	a :="aaa"
	ssh := *(*reflect.StringHeader)(unsafe.Pointer(&a))
	b := *(*[]byte)(unsafe.Pointer(&ssh))
	fmt.Printf("%v",b)
}

func Test_memory_copy1(t *testing.T) {
	a :="aaa"
	ssh := *(*reflect.StringHeader)(unsafe.Pointer(&a))
	b := *(*reflect.SliceHeader)(unsafe.Pointer(&ssh))
	fmt.Printf("%v",b)
}
