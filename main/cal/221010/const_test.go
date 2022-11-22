package _21010

import (
	"fmt"
	"testing"
)

const (
	a = iota
	b
)

const (
	name = "dwliang"
	name1 = "dwliang"
	c = iota
	d = iota
)

type Student1 struct {
	Name string
}

func Test_const(t *testing.T) {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println(&Student1{Name: "menglu"} == &Student1{Name: "menglu"})
	fmt.Println(Student1{Name: "menglu"} == Student1{Name: "menglu"})

	fmt.Println([...]string{"1"} == [...]string{"1"})
	//fmt.Println([]string{"1"} == []string{"1"})

	kv := map[string]*Student1{"menglu": {Name: "21"}}
	kv["menglu"].Name = "22"
	s := []Student1{{Name: "21"}}
	s[0].Name = "22"
	fmt.Println(kv["menglu"], s)
}
