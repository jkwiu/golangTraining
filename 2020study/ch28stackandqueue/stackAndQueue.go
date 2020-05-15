package ch28stackandqueue

import (
	"fmt"
)

type Node struct {
	Next *Node
	Prev *Node
	Val  int
}

type List struct {
	Root *Node
	Tail *Node
}

type Stack struct {
	l *List
}

type Queue struct {
	l *List
}

func (l *List) AddNode(val int) {
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

func (l *List) RemoveList(n *Node) {
	if l.Root == nil {
		fmt.Println("엄써요~")
		return
	}

	// 맨 앞
	if n == l.Root {
		root := l.Root.Next
		root.Prev = nil
		l.Root = root
	} else if n == l.Tail { // 맨 끝
		tail := l.Tail.Prev
		l.Tail.Prev = nil
		tail.Next = nil
		l.Tail = tail
	} else { // 중간
		n.Prev.Next = n.Next
		n.Next.Prev = n.Prev
	}
}

func (l *List) PrintNodes() {
	node := l.Root
	for node.Next != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}
	fmt.Println(node.Val)
}

func NewStack() *Stack {
	return &Stack{l: &List{}}
}

func (s *Stack) Push(val int) {
	s.l.AddNode(val)
}

func (s *Stack) Pop() int {
	rst := s.l.Tail.Val
	s.l.RemoveList(s.l.Tail)
	return rst
}

func (s *Stack) PrintStacks() {
	node := s.l.Tail
	for node.Prev != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Prev
	}
	fmt.Println(node.Val)
}

func NewQueue() *Queue {
	return &Queue{l: &List{}}
}

func (q *Queue) Push(val int) {
	q.l.AddNode(val)
}

func (q *Queue) Pop() int {
	rst := q.l.Root.Val
	q.l.RemoveList(q.l.Root)
	return rst
}

func (q *Queue) PrintQueues() {
	node := q.l.Root
	for node.Next != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}
	fmt.Println(node.Val)
}
