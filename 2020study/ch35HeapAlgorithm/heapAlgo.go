// package ch35HeapAlgorithm
package main

import "fmt"

func main() {
	h := &Heap{}
	nums := []int{-1, 3, -1, 5, 4}

	for i := 0; i < len(nums); i++ {
		h.Push(nums[i])
		if h.Count() > 1 {
			h.Pop()
		}
	}
	fmt.Println("1번째 큰 값: ", h.Pop())

	h = &Heap{}

	for i := 0; i < len(nums); i++ {
		h.Push(nums[i])
		if h.Count() > 2 {
			h.Pop()
		}
	}
	fmt.Println("2번째 큰 값: ", h.Pop())

	h = &Heap{}

	for i := 0; i < len(nums); i++ {
		h.Push(nums[i])
		if h.Count() > 3 {
			h.Pop()
		}
	}
	fmt.Println("3번째 큰 값: ", h.Pop())

	h = &Heap{}

	for i := 0; i < len(nums); i++ {
		h.Push(nums[i])
		if h.Count() > 4 {
			h.Pop()
		}
	}
	fmt.Println("4번째 큰 값: ", h.Pop())

	h = &Heap{}

	for i := 0; i < len(nums); i++ {
		h.Push(nums[i])
		if h.Count() > 5 {
			h.Pop()
		}
	}
	fmt.Println("5번째 큰 값: ", h.Pop())
}

// 최소힙
type Heap struct {
	list []int
}

func (h *Heap) Push(v int) {
	h.list = append(h.list, v)
	idx := len(h.list) - 1
	for idx >= 0 {
		parentIdx := (idx - 1) / 2
		if parentIdx < 0 {
			break
		}
		if h.list[idx] < h.list[parentIdx] {
			h.list[idx], h.list[parentIdx] = h.list[parentIdx], h.list[idx]
			idx = parentIdx
		} else {
			break
		}
	}
}

func (h *Heap) Pop() int {
	if len(h.list) == 0 {
		return 0
	}
	top := h.list[0]
	last := h.list[len(h.list)-1]
	h.list = h.list[:len(h.list)-1]

	if len(h.list) == 0 {
		return top
	}
	h.list[0] = last

	idx := 0
	for idx < len(h.list) {
		swapIdx := -1
		leftIdx := (2 * idx) + 1
		if leftIdx >= len(h.list) {
			break
		}
		if h.list[leftIdx] < h.list[idx] {
			swapIdx = leftIdx
		}
		rightIdx := (2 * idx) + 2
		if rightIdx < len(h.list) {
			if h.list[rightIdx] < h.list[idx] {
				if swapIdx < 0 || h.list[rightIdx] < h.list[swapIdx] {
					swapIdx = rightIdx
				}
			}
		}

		if swapIdx < 0 {
			break
		}

		h.list[idx], h.list[swapIdx] = h.list[swapIdx], h.list[idx]
		idx = swapIdx
	}
	return top
}

func (h *Heap) Print() {
	fmt.Println(h.list)
}

func (h *Heap) Count() int {
	return len(h.list)
}
