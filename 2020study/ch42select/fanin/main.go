package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	c1, c2, c3 := make(chan int), make(chan int), make(chan int)
	sendInts := func(c chan<- int, begin, end int) {
		defer close(c)

		for i := begin; i < end; i++ {
			c <- i
		}
	}
	go sendInts(c1, 11, 13)
	go sendInts(c2, 21, 23)
	go sendInts(c3, 31, 40)

	for n := range FanIn3(c1, c2, c3) {
		fmt.Println("NumGoroutine: ", runtime.NumGoroutine())
		fmt.Println(n, ",")
	}
}

func FanIn3(in1, in2, in3 <-chan int) <-chan int {
	out := make(chan int)
	openCnt := 3
	closeChan := func(c *<-chan int) bool {
		*c = nil
		openCnt--
		return openCnt == 0
	}
	go func() {
		defer close(out)
		for {
			select {
			case n, ok := <-in1:
				if ok {
					out <- n
				} else if closeChan(&in1) {
					return
				}
			case n, ok := <-in2:
				if ok {
					out <- n
				} else if closeChan(&in2) {
					return
				}
			case n, ok := <-in3:
				if ok {
					out <- n
				} else if closeChan(&in3) {
					return
				}
			// 채널과의 통신을 5초간 기다리고 안오면 이 함수를 빠져나가게 구성한다.
			case <-time.After(5 * time.Second):
				fmt.Println("No send and receive communication for 5 seconds")
				return
			default:
				fmt.Println("Data is not ready. Skipping")
			}
		}
	}()
	return out
}
