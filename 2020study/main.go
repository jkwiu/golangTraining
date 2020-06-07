package main

import (
	"fmt"

	"github.com/golangTraining/2020study/tdd"
)

func main() {
	var h *tdd.Heap
	nums := []int{-1, 3, -1, 5, 4}

	// 1번째 작은 수
	h = &tdd.Heap{}
	for i := 0; i < len(nums); i++ {
		h.Push(nums[i])
		if h.Count() > 1 {
			h.Pop()
		}
	}
	h.Print()
	fmt.Println(h.Pop())

	// 2번째 작은 수
	h = &tdd.Heap{}
	for i := 0; i < len(nums); i++ {
		h.Push(nums[i])
		if h.Count() > 2 {
			h.Pop()
		}
	}
	h.Print()
	fmt.Println(h.Pop())

	// 3번째 작은 수
	h = &tdd.Heap{}
	for i := 0; i < len(nums); i++ {
		h.Push(nums[i])
		if h.Count() > 3 {
			h.Pop()
		}
	}
	h.Print()
	fmt.Println(h.Pop())

}
