package mianshi

import (
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"testing"
	"time"
)

// 這題考量for range會創建每個變量的副本而不是引用
func Test_foreach(t *testing.T) {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		m[key] = &val
		//value := val
		//m[key] = &value  正确的使用方法
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}

// 考察slice的声明方法
func Test_append(t *testing.T) {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)

	s1 := make([]int, 0)
	s1 = append(s1, 1, 2, 3, 4)
	fmt.Println(s1)

	s2 := make([]int, 0, 5)
	s2 = append(s2, 1, 2, 3, 4)
	fmt.Println(s2)
}

// make和new的区别 new是创建一个初始值为“nil”的指针
func Test_makeAndNew(t *testing.T) {
	m := make(map[int]string)
	m[1] = "test"

	m1 := new(map[int]string)
	tmp := *m1
	tmp[1] = "dwliang"
}

func Test_sliceAppend(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	//s1 = append(s1, s2) 错误写法
	s1 = append(s1, s2...) // 正确写法
	fmt.Println(s1)
}

// struct的比较 不仅需要各个属性类型可比较 也要确保字段声明的顺序  顺序不同则不同
func Test_compareStruct(t *testing.T) {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	//sn11 := struct {
	//	name string
	//	age  int
	//
	//}{age: 11, name: "qq"}
	//
	//sn21 := struct {
	//	age  int
	//	name string
	//}{age: 11, name: "qq"}
	//
	//if sn11 == sn21 {
	//	fmt.Println("sn1 == sn2")
	//}

	//sm1 := struct {
	//	age int
	//	m   map[string]string
	//}{age: 11, m: map[string]string{"a": "1"}}
	//
	//sm2 := struct {
	//	age int
	//	m   map[string]string
	//}{age: 11, m: map[string]string{"a": "1"}}
	//
	////map类型不能进行比较
	//if sm1 == sm2 {
	//	fmt.Println("sm1 == sm2")
	//}
}

func Test_AttributeAccess(t *testing.T) {
	s := &struct {
		id   int
		name string
	}{id: 1, name: "dwliang"}

	fmt.Println(s.name)
	fmt.Println((*s).name)
}

type MyInt1 int
type MyInt2 = int

func Test_typeDeclare(t *testing.T) {
	var i int = 0
	var j MyInt1 = 1
	var i1 MyInt1 = j
	var i2 MyInt2 = i
	fmt.Println(i1, i2)
}

func Test_Abstract(t *testing.T) {
	//var s1 []int
	c := Work{3}
	var a A = c
	var b B = c
	fmt.Println(a.ShowA())
	fmt.Println(b.ShowB())
}

func Test_map(t *testing.T) {
	var m = make(map[string]int) //A
	m["a"] = 1
	if v := m["b"]; v != 0 { //B
		fmt.Println(v)
	}
}

func incr(p *int) int {
	*p++
	return *p
}

type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w Work) ShowA() int {
	return w.i + 10
}

func (w Work) ShowB() int {
	return w.i + 20
}

func testDuration(t time.Duration) {
	fmt.Println(int(t / time.Millisecond))
}

//func genRandNums(min, max float64) []float64 {
//	var randNums []float64
//	s := rand.NewSource(time.Now().Unix())
//	r := rand.New(s)
//	for x := 0; x < 10; x++ {
//		// generate random float in range of min and max inclusive, append
//		// to randNums and return randNums
//		randNums[i] = min + rand.Float64()*(max-min)
//	}
//	return randNums
//}

func randFloats(min, max float64, n int) []float64 {
	res := make([]float64, n)

	for i := range res {
		res[i] = min + rand.Float64()*(max-min)
	}
	return res
}

func Test_test1(t *testing.T) {
	nums := randFloats(0.5, 2.3, 1)
	fmt.Println(nums[0])

	testDuration(5000 * time.Millisecond)
	mp := sync.Map{}

	actual, loaded := mp.LoadOrStore("d", "")
	fmt.Println(actual)
	fmt.Println(loaded)

	fmt.Println("--------------")

	actual1, loaded1 := mp.LoadOrStore("d", "")
	fmt.Println(actual1)
	fmt.Println(loaded1)

	testDuration(time.Duration(10) * time.Second)
	now := time.Now()
	deadline := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	fmt.Println(deadline)
	timestamp1, _ := time.Parse("2006-01-02 15:04:05", "2020-01-16 18:48:00")
	fmt.Println(timestamp1)

	i := 1
	of := reflect.ValueOf(&i)
	of.Elem().SetInt(11)
	fmt.Println(i)
}

func Test_slipwindow(t *testing.T) {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

func maxSlidingWindow(nums []int, k int) []int {
	var max []int
	if len(nums) < 1 {
		return max
	}
	var dq []int
	for i := 0; i < len(nums); i++ {
		if len(dq) == 0 {
			dq = append(dq, i)
		} else if nums[i] > nums[dq[0]] {
			dq = []int{i}
		} else if i-dq[0] > k-1 {
			dq = dq[1:]
			j := len(dq) - 1
			for ; (j >= 0) && (nums[i] >= nums[dq[j]]); j-- {

			}
			if j < 0 {
				dq = []int{i}
			} else {
				dq = append(dq[0:j+1], i)
			}
		} else {
			j := len(dq) - 1
			for ; (j > 0) && (nums[i] >= nums[dq[j]]); j-- {

			}
			dq = append(dq[0:j+1], i)
		}
		fmt.Println(dq)
		if i >= k-1 {
			max = append(max, nums[dq[0]])
		}
	}
	return max
}
