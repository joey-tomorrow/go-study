package _21010

import (
	"fmt"
	"sync"
	"testing"
)

/*
使⽤两个 goroutine 交替打印序列，⼀个goroutine 打印数字， 另外⼀
个 goroutine 打印字⺟， 最终效果如下：
12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728

解题思路
问题很简单，使⽤ channel 来控制打印的进度。使⽤两个 channel ，来分别控制数字和
字⺟的打印序列， 数字打印完成后通过 channel 通知字⺟打印, 字⺟打印完成后通知数
字打印，然后周⽽复始的⼯作。
源码参考
*/
func Test_jiaotidayin(t *testing.T) {
	var (
		charChan,numChan = make(chan bool), make(chan bool)

		wait = sync.WaitGroup{}
	)

	go func() {
		i := 0
		for  {
			select {
			case <-numChan:
				i ++
				fmt.Print(i)
				i ++
				fmt.Print(i)
				charChan <- true
				break
			default:
				break
			}

		}
	}()

	wait.Add(1)

	go func(wait *sync.WaitGroup) {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		for {

			select {
			case <-charChan:
				if i >= len(str) - 1 {
					wait.Done()
					return
				}
				fmt.Print(str[i:i+1])
				i ++
				if i >= len(str) {
					wait.Done()
					return
				}
				fmt.Print(str[i:i+1])
				i ++
				numChan <- true
				break
			default:
				break
			}
		}
	}(&wait)

	numChan <- true
	wait.Wait()
}

func Test_jiaotidayin1(t *testing.T) {
	var (
		numChan, charChan = make(chan int), make(chan string)
		waitGroup = sync.WaitGroup{}
	)

	go func() {
		i := 1
		for {
			select {
				case <- numChan:
					fmt.Print(i)
					i++
					fmt.Print(i)
					i++
					charChan <- ""
					break
			default:
				break
			}
		}

		fmt.Println("end num")
	}()

	waitGroup.Add(1)

	go func(wait *sync.WaitGroup) {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		for{
			select {
			case <- charChan:
				if i >= len(str) {
					wait.Done()
					break
				}
				fmt.Print(string(str[i]))
				i ++
				if i >= len(str) {
					wait.Done()
					break
				}
				fmt.Print(string(str[i]))
				i ++
				numChan <- 0
				break
			default:
				break
			}
		}
		fmt.Println("end char")
	}(&waitGroup)

	numChan <- 0
	waitGroup.Wait()
}
