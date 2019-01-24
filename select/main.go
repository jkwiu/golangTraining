//어떤 일을 하고 있다가, 다른 입력이 들어올 경우 다른 일을 할 수 있도록 해줌

package main

import (
	"fmt"
	"time"
)

func pop(c chan int) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		c <- i
		i++
	}
}

func main() {
	c := make(chan int)
	go pop(c)
	//특정시간 이후에 실행 after channel
	timerChan := time.After(10 * time.Second)
	//일정시간마다 실행 tick channel
	tickTimerChan := time.Tick(2 * time.Second)

	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-timerChan:
			fmt.Println("TimeOut")
			return
		case <-tickTimerChan:
			fmt.Println("Tick")
			//default:
			//	fmt.Println("Idle")
			//	time.Sleep(1 * time.Second)
		}
	}
}
