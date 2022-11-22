package _21010

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type Admin struct {
	Name string
	Age int
}

func Test_uintptr(t *testing.T) {
	admin := Admin{
		Name: "seekload",
		Age:  18,
	}
	ptr := &admin
	fmt.Println(*ptr) //输出:{seekload 18}
	//这样的方式的话,获取的是第一个字段的指针
	name := (*string)(unsafe.Pointer(ptr))
	fmt.Println(name)
	*name = "四哥"
	fmt.Println(*ptr) //输出:{四哥 18}
	/* 成员变量 Age 不是第一个字段，想要修改它的值就需要内存偏移。我们先将 admin 的指针转化为 uintptr，
	再通过 unsafe.Offsetof() 获取到 Age 的偏移量，两者都是 uintptr，进行相加指针运算获取到成员 Age 的地址，
	最后需要将 uintptr 转化为 unsafe.Pointer，再转化为 *int，才能对 Age 操作。*/
	age := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + unsafe.Offsetof(ptr.Age))) // 2
	fmt.Println(reflect.TypeOf(age))                                                       //输出:int
	*age = 35
	fmt.Println(*ptr) //输出:{四哥 35}

}
