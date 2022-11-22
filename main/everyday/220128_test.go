package everyday

import (
	"fmt"
	"testing"
)



func Test_220128(t *testing.T)  {
	a := []int{1,2,3,4}

	//遍历切片
	for k,v  := range a{
		fmt.Println(k,v)
		fmt.Println(&k,"==",&v)  //地址相同
	}

	//遍历切片
	for k,v  := range a{
		kk,vv := k,v   //new一个变量
		fmt.Println(k,v)
		fmt.Println(&kk,"==",&vv)  //地址不同，达到要求
	}
}
