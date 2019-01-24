package dataStruct

import "fmt"

type Heap struct {
	list []int
}

func (h *Heap) Push(v int) {
	//v를 맨 뒤에 추가한다.
	h.list = append(h.list, v)

	idx := len(h.list) - 1

	//부모로 올라가면서 비교
	for idx >= 0 {
		parentIdx := (idx - 1) / 2
		if parentIdx < 0 {
			break
		}

		//크기를 비교해서 swap
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

	//heap이 아예 없는 경우
	if len(h.list) == 0 {
		return 0
	}

	//heap tree의 가장 위와 가장 아래를 pop
	//마지막을 잘라낸다.
	//맨 뒤의 애를 맨 위로 바꾼다.
	top := h.list[0]
	last := h.list[len(h.list)-1]
	h.list = h.list[:len(h.list)-1]

	//자식노드로 가면서 자신보다 큰 애와 swap
	h.list[0] = last
	idx := 0
	for idx < len(h.list) {
		leftIdx := idx*2 + 1
		//swapIdx = 바꾼 idx
		swapIdx := -1

		//자식노드가 없으면
		if leftIdx >= len(h.list) {
			break
		}
		if h.list[leftIdx] > h.list[idx] {
			swapIdx = leftIdx
		}

		rightIdx := idx*2 + 2
		if rightIdx < len(h.list) {
			if h.list[rightIdx] > h.list[idx] {
				if swapIdx < 0 || h.list[swapIdx] < h.list[rightIdx] {
					swapIdx = rightIdx
				}
			}
		}

		//바꿀애가 없으면 break
		if swapIdx < 0 {
			break
		}
		h.list[idx], h.list[swapIdx] = h.list[swapIdx], h.list[idx]
		idx = swapIdx
	}
	return top
}
