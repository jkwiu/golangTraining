package main

import (
	"sync"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)
	c := FanIn(c1, c2, c3)

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		c1 <- 1
		c2 <- 2
		c3 <- 3
	}()

	go func() {
		wg.Wait()
		close(c1)
		close(c2)
		close(c3)
	}()

	println(<-c)
	println(<-c)
	println(<-c)
}

func FanIn(ins ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(ins))
	for _, in := range ins {
		go func(in <-chan int) {
			defer wg.Done()
			for num := range in {
				out <- num
			}
		}(in)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
