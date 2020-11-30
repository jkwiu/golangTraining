package main

import (
	"fmt"
	"sync"
)

func main() {
	initialize2()
}

func initialize1() {
	done := make(chan struct{})
	go func() {
		// 초기화를 한번만 실행
		defer close(done)
		fmt.Println("Initialized")
	}()
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-done
			fmt.Println("Goroutine: ", i)
		}(i)
	}
	wg.Wait()
}

// 이렇게 초기화코드를 사용할 것을 권장한다.
func initialize2() {
	var once sync.Once
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			once.Do(func() {
				fmt.Println("Initilized")
			})
			fmt.Println("Goroutine: ", i)
		}(i)
	}
	wg.Wait()
}
