package ch27packaging

import "fmt"

type Node struct {
	Next *Node
	Prev *Node
	Val int
}

type List struct {
	Root *Node
	Tail *Node
}

// 리스트 추가
func (l *List) AddList(val int){
	// 노드가 없을 때 리스트 생성
	if l.Root == nil {
		l.Root = &Node{}
		l.Tail = l.Root
		fmt.Println("리스트의 노드가 생성되었습니다.")
		return
	}

	// 꼬리에다 새 ele 추가
	l.Tail.Next = &Node{Val: val}
	l.Tail.Next.Prev = l.Tail
	l.Tail = l.Tail.Next
	fmt.Println("리스트 ele가 추가되었습니다.")
}

// 리스트 삭제
func (l *List) RemoveList(n *Node) {
	// 예외 처리
	// 리스트가 없을 때
	if l.Root == nil {
		fmt.Println("리스트가 없습니다.")
	}

	// 맨 처음 삭제
	if l.Root == n {
		l.Root = l.Root.Next
		l.Root.Prev = nil
		n.Next = nil
	} else if l.Tail == n {	// 맨 끝 삭제
		l.Tail = l.Tail.Prev
		n.Prev = nil
		l.Tail.Next = nil
	} else {	// 중간 삭제
		n.Prev.Next = n.Next
		n.Next.Prev = n.Prev
	}	
}

// 리스트 출력
func (l *List) PrintList() {
	node := l.Root

	for node.Next != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}
	fmt.Printf("%d\n", node.Val)
}