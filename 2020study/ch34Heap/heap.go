package ch34Heap

import "fmt"

type Heap struct {
	list []int
}

func (h *Heap) Push(v int) {
	// 맨 아래서부터 위로 올라가면서 크기 비교
	h.list = append(h.list, v)

	idx := len(h.list) - 1
	for idx >= 0 {
		parentIdx := (idx - 1) / 2
		if parentIdx < 0 {
			break
		}
		if h.list[idx] > h.list[parentIdx] {
			h.list[idx], h.list[parentIdx] = h.list[parentIdx], h.list[idx]
			idx = parentIdx
		} else {
			break
		}
	}
}

func (h *Heap) Print() {
	fmt.Println(h.list)
}

func (h *Heap) Pop() int {
	// 맨 위부터 시작해서 아래로 가면서 크기 비교
	// list가 비어있는 경우 err출력대신에 임시로 0
	if len(h.list) == 0 {
		fmt.Println("빈 배열입니다.")
		return 0
	}

	// 맨 위의 값을 저장
	top := h.list[0]
	// 맨 마지막 값을 맨 위로
	last := h.list[len(h.list)-1]
	if len(h.list) != 1 {
		h.list = h.list[:len(h.list)-1]
		h.list[0] = last
	} else { // 배열의 길이가 1일 때는
		h.list = h.list[:len(h.list)-1]
	}

	// 맨 위부터 시작해서 자식노드와 크기를 비교
	idx := 0
	for idx < len(h.list) { // 맨 마지막 idx가 아닌 이상 loop
		swapIdx := -1        // swap할 idx표시
		leftIdx := idx*2 + 1 // 왼쪽 자식 노드

		// leftidx가 없는 경우
		if leftIdx >= len(h.list) {
			break
		}

		// leftIdx가 부모보다 큰 경우
		if h.list[leftIdx] > h.list[idx] {
			swapIdx = leftIdx
		}

		// rightIdx 비교
		rightIdx := idx*2 + 2
		if rightIdx < len(h.list) {
			if h.list[rightIdx] > h.list[idx] {
				if swapIdx < 0 || h.list[swapIdx] < h.list[rightIdx] {
					swapIdx = rightIdx
				}
			}
		}

		// 바꿀 애가 없다리...
		if swapIdx < 0 {
			break
		}
		h.list[idx], h.list[swapIdx] = h.list[swapIdx], h.list[idx]
		idx = swapIdx
	}
	return top
}
