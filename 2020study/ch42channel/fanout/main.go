package main

import "time"

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			for n := range c {
				time.Sleep(1)
				println(i, n)
			}
		}
	}()
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}
