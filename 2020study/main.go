package main

import (
	"golangTraining/2020study/ch28stackandqueue"
)

func main() {
	queue := ch28stackandqueue.NewQueue()

	for i := 0; i < 10; i++ {
		queue.Push(i)
	}

	queue.PrintQueues()

}
