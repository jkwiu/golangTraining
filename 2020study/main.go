package main

import (
	"fmt"
	"golangTraining/2020study/ch28stackandqueue"
)

func main() {
	s := &ch28stackandqueue.Stack{}
	s.Push(s)
	fmt.Println(s)
}
