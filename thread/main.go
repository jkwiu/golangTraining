package main

import (
	"fmt"
	"time"
)

func main() {
	//go thread로 fun1()을 수행하라.
	go fun1()
	for i := 0; i < 20; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("main", i)
	}
	//한 줄 입력하는 것 기다리는 것
	fmt.Scanln()
}

func fun1() {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("fun1:", i)
	}
}
