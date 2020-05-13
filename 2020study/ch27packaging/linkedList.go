package ch27packaging

import "fmt"

type Node struct {
	Prev *Node
	Next *Node
	Val  int
}

type DoubleLinkedList struct {
	Root *Node
	Tail *Node
}

func (l *DoubleLinkedList) AddNode(val int) {
	// node가 없을 때 하나를 생성
	if l.Root == nil {
		l.Root = &Node{Val: val}
		l.Tail = l.Root
		return
	}
	l.Tail.Next = &Node{Val: val}
	prev := l.Tail
	l.Tail = l.Tail.Next
	l.Tail.Prev = prev
}

func (l *DoubleLinkedList) RemoveNode(node *Node) {
	if node == l.Root {
		l.Root = l.Root.Next
		if l.Root != nil {
			l.Root.Prev = nil
		}
		node.Next = nil
		return
	}

	prev := node.Prev

	if node == l.Tail {
		prev.Next = nil
		l.Tail.Prev = nil
		l.Tail = prev
	} else { // 중간 노드를 지울 때
		node.Prev = nil
		prev.Next = prev.Next.Next
		prev.Next.Prev = prev
	}
	node.Next = nil
}

func (l *DoubleLinkedList) PrintNodes() {
	node := l.Root
	for node.Next != nil {
		fmt.Printf("%d ->", node.Val)
		node = node.Next
	}
	fmt.Printf("%d\n", node.Val)
}

func (l *DoubleLinkedList) PrintReverse() {
	node := l.Tail
	for node.Prev != nil {
		fmt.Printf("%d ->", node.Val)
		node = node.Prev
	}
	fmt.Printf("%d\n", node.Val)
}

// 맨 뒤의 값을 반환하는 함수
func (l *DoubleLinkedList) Back() int {
	if l.Tail != nil {
		return l.Tail.Val
	}
	return 0
}

// 맨 뒤의 값을 지우는 함수
func (l *DoubleLinkedList) PopBack() {
	if l.Tail == nil {
		return
	}
	l.RemoveNode(l.Tail)
}

func (l *DoubleLinkedList) Empty() bool {
	return l.Root == nil
}

func (l *DoubleLinkedList) Front() int {
	if l.Root != nil {
		return l.Root.Val
	}
	return 0
}

// 맨 앞의 것을 없앤다
func (l *DoubleLinkedList) PopFront() {
	if l.Root == nil {
		return
	}
	l.RemoveNode(l.Root)
}
