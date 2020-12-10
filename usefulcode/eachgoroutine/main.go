package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	inje := &WPPIS{Name: "inje", UpdateRate: 1000}
	kangwon := &WPPIS{Name: "kangwon", UpdateRate: 3000}
	jeju := &WPPIS{Name: "jeju", UpdateRate: 3000}
	daejeon := &WPPIS{Name: "daejeon", UpdateRate: 5000}
	seoul := &WPPIS{Name: "seoul", UpdateRate: 1000}

	for {
		// rstChan := make(chan string, 200)
		wg.Add(5)
		go inje.start()
		go kangwon.start()
		go jeju.start()
		go daejeon.start()
		go seoul.start()
		wg.Wait()
	}
}

type WPPIS struct {
	Name       string
	UpdateRate int
}

func (w *WPPIS) start() {
	start := time.Now()
	for {
		rstChan := make(chan string, 1000)
		defer wg.Done()
		rstChan <- "I am :" + w.Name + ", Update rate is: " + strconv.Itoa(w.UpdateRate) + ", 시작한지는: " + time.Since(start).String() + ", Go routine num는 " + strconv.Itoa(runtime.NumGoroutine())
		go func(rstChan chan string) {
			for ele := range rstChan {
				fmt.Println(ele)
			}
		}(rstChan)
		close(rstChan)
		time.Sleep(time.Duration(w.UpdateRate) * time.Millisecond)
	}
}

func (w *WPPIS) stop() {

}
