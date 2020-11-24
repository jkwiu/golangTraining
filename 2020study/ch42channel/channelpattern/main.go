package main

import "fmt"

func main() {
	// Example_simpleChannel()
	PlusTwo := Chain(PlusOne, PlusOne)
	c := make(chan int)
	go func() {
		defer close(c)
		c <- 5

	}()
	for num := range PlusTwo(c) {
		println(num)
	}
}

// 자주 쓰이는 패턴
// 함수가 채널을 반환하게 만드는 패턴
func Example_simpleChannel() {
	c := func() <-chan int {
		c := make(chan int)
		go func() {
			defer close(c)
			c <- 1
			c <- 2
			c <- 3
		}()
		return c
	}()
	for num := range c {
		fmt.Println(num)
	}
}

func PlusOne(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			out <- num + 1
		}
	}()
	return out
}

func ExamplePlusOne() {
	c := make(chan int)
	go func() {
		defer close(c)
		c <- 5
		c <- 3
		c <- 8
	}()
	for num := range PlusOne(PlusOne(c)) {
		println(num)
	}
}

type IntPipe func(<-chan int) <-chan int

func Chain(ps ...IntPipe) IntPipe {
	return func(in <-chan int) <-chan int {
		c := in
		for _, p := range ps {
			c = p(c)
		}
		return c
	}
}
