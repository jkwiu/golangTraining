package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	asyncWait()
}

// race condition err가 뜬다
func raceConditonErr() {
	cnt := int64(10)
	for i := 0; i < 10; i++ {
		go func() {
			cnt-- // 이 부분과
		}()
	}
	for cnt > 0 { // 이 부분
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println(cnt)
}

// race condition err가 안뜬다
func noRaceConditonErr() {
	cnt := int64(10)
	for i := 0; i < 10; i++ {
		go func() {
			atomic.AddInt64(&cnt, -1) // 이 부분
		}()
	}
	for atomic.LoadInt64(&cnt) > 0 { // 이 부분
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println(cnt)
}

// race condition 안뜸
func usefulChannel() {
	req, resp := make(chan struct{}), make(chan int64)
	cnt := int64(10)
	go func(cnt int64) {
		defer close(resp)
		for _ = range req {
			cnt--
			resp <- cnt
		}
	}(cnt)
	for i := 0; i < 10; i++ {
		go func() {
			req <- struct{}{}
		}()
	}
	for cnt = <-resp; cnt > 0; cnt = <-resp {

	}
	close(req)
	fmt.Println(cnt)
}

// 보기 깔끔
func asyncWait() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
		}()
	}
	wg.Wait()
}
