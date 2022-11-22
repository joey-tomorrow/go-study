package _21010

import (
	"fmt"
	"runtime"
	"testing"
)

func Test_sixunhuan(t *testing.T)  {
	var i byte
	go func() {
		for i = 0; i <= 255; i++ {
		}

		for{}
	}()

	fmt.Println("Dropping mic")
	// Yield execution to force executing other goroutines
	runtime.Gosched()
	runtime.GC()
	fmt.Println("Done")
	//time.Sleep(30 * time.Second)


	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)

	var peo People1 = &Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))

	if live() == nil {
		fmt.Println("aaaaa")
	} else {
		fmt.Println("bbbbbb")
	}

	var stu *Student
	fmt.Println(stu == nil)
}

type People1 interface {
	Speak(string) string
}

type Student struct{}
func (stu *Student) Speak(think string) (talk string) {
	return think
}

func live() People1 {
	var stu *Student
	return stu
}