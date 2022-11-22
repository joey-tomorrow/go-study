package _21010

import (
	"fmt"
	"sync"
	"testing"
	"time"
)


/*
解析：
看到阻塞协程第⼀个想到的就是 channel ，题⽬中要求并发安全，那么必须⽤锁，还要
实现多个 goroutine 读的时候如果值不存在则阻塞，直到写⼊值，那么每个键值需要有
⼀个阻塞 goroutine 的  channel 。
实现如下：
⾼并发下的锁与map的读写 场景：在⼀个⾼并发的web服务器中，要限制IP的频繁访问。现模拟100个IP同时并发访问服 务器，每个IP要重复访问1000次。 每个IP三分钟之内只能访问⼀次。修改以下代码完成该过程，要求能成功输出 success:100 Out(key string, val interface{}) //存入key /val，如果该key读取的
goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
*/
type SMap interface {
	Out(key string, val interface{}) //存入key /val，如果该key读取的 goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
	Rd(key string, timeout time.Duration) interface{} //读取一个key，如果 key不存在阻塞，等待key存在或者超时
}

type SyncMapTest struct {
	mu sync.RWMutex
	Element map[string]*SMValue
}

type SMValue struct {
	Value interface{}
	wg sync.WaitGroup
	isExist bool
}

func (s *SyncMapTest) Out(key string, val interface{}) {
	s.mu.Lock()
	if s.Element == nil {
		s.Element = make(map[string]*SMValue)
		s.Element[key] = &SMValue{isExist: true, Value: val}
		s.mu.Unlock()
		return
	}

	value, ok := s.Element[key]
	if !ok {
		s.Element[key] = &SMValue{isExist: true, Value: val}
		s.mu.Unlock()
	}
	if ok {
		value.Value = val
		value.isExist = true
		value.wg.Done()
		s.mu.Unlock()
	}
}

func (s *SyncMapTest) Rd(key string, timeout time.Duration) interface{} {
	s.mu.RLock()
	if s.Element == nil {
		s.Element = make(map[string]*SMValue)
	}

	value, ok := s.Element[key]
	if !ok {
		value = &SMValue{
			isExist: false,
		}
		s.Element[key] = value
		value.wg.Add(1)
	}

	if value.isExist {
		s.mu.RUnlock()
		return value.Value
	}

	s.mu.RUnlock()

	//go func() {
	//	timer := time.NewTimer(timeout)
	//	defer timer.Stop()
	//
	//	select {
	//	case <-timer.C:
	//		value.wg.Done()
	//	}
	//}()
	//
	//value.wg.Wait()

	WaitTimeout(&value.wg, timeout)

	if value.isExist {
		return value.Value
	}

	fmt.Println("not exists")
	return nil
}


func Test_syncMap(t *testing.T) {
	mapT := &SyncMapTest{}

	go func() {
		time.Sleep(2 * time.Second)
		mapT.Out("dwliang", "test")
	}()

	go func() {
		fmt.Println(mapT.Rd("dwliang", 3 * time.Second))
	}()

	go func() {
		fmt.Println(mapT.Rd("dwliang", 3 * time.Second))
	}()

	time.Sleep(4 * time.Second)
}

func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	// 要求手写代码
	// 要求sync.WaitGroup支持timeout功能
	// 如果timeout到了超时时间返回true
	// 如果WaitGroup自然结束返回false
	c := make(chan bool)

	go time.AfterFunc(timeout, func() {
		c <- true
	})

	go func() {
		wg.Wait()
		c <- false
	}()

	return <- c
}

func TestMapSync(t *testing.T) {
	var m sync.Map
	m.LoadOrStore("a", 1)
	m.Delete("a")
	//fmt.Println(len(m))
}