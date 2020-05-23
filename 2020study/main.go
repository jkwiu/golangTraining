package main

import (
	"fmt"
	"golangTraining/2020study/ch34Heap"
)

func main() {
	h := &ch34Heap.Heap{}

	h.Push(6)
	h.Push(6)
	h.Push(7)
	h.Push(8)
	h.Push(5)
	h.Push(9)

	h.Print()

	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
}
