package tdd

import "fmt"

// 1. linked list
type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

type List struct {
	root *Node
	tail *Node
}

func NewList(val int) *List {
	list := &List{root: &Node{Val: val}}
	list.tail = list.root
	return list
}

func (l *List) AddNode(val int) int {
	if l.root == nil {
		fmt.Println("There is no List!!")
		return 0
	}
	l.tail.Next = &Node{Val: val}
	l.tail.Next.Prev = l.tail
	l.tail = l.tail.Next
	return l.tail.Val
}

func (l *List) RemoveNode(n *Node) int {
	if l.root == nil {
		fmt.Println("There is no List!!")
		return 0
	}
	rst := n.Val
	if n == l.root { // 맨 앞일 때
		l.root = l.root.Next
		l.root.Prev.Next = nil
		l.root.Prev = nil
	}
	return rst
}

func (l *List) PrintNodes() {
	node := l.root
	for node.Next != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}
	fmt.Println(node.Val)
}

// stack
type Stack struct {
	list *List
}

// 신규
func NewStack(val int) *Stack {
	stack := &Stack{list: &List{root: &Node{Val: val}}}
	stack.list.tail = stack.list.root
	return stack
}

// 추가
func (s *Stack) Push(val int) int {
	return s.list.AddNode(val)
}

// 삭제
func (s *Stack) Pop() int {
	return s.list.RemoveNode(s.list.tail)
}

// 출력
func (s *Stack) PrintStack() {
	node := s.list.tail
	for node.Prev != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Prev
	}
	fmt.Println(node.Val)
}

// queue
type Queue struct {
	list *List
}

// 신규
func NewQueue(val int) *Queue {
	queue := &Queue{list: &List{root: &Node{Val: val}}}
	queue.list.tail = queue.list.root
	return queue
}

// 추가
func (q *Queue) Push(val int) int {
	return q.list.AddNode(val)
}

// 삭제
func (q *Queue) Pop() int {
	return q.list.RemoveNode(q.list.root)
}

// 출력
func (q *Queue) PrintQueue() {
	q.list.PrintNodes()
}
