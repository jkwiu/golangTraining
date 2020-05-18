package main

import "fmt"

type Node struct {
	prev *Node
	next *Node
	val  int
}

type DoubleLinkedList struct {
	root *Node
	tail *Node
}

func (l *DoubleLinkedList) AddNode(val int) {
	// node가 없을 때 하나를 생성
	if l.root == nil {
		l.root = &Node{val: val}
		l.tail = l.root
		return
	}
	l.tail.next = &Node{val: val}
	prev := l.tail
	l.tail = l.tail.next
	l.tail.prev = prev
}

func (l *DoubleLinkedList) RemoveNode(node *Node) {
	if node == l.root {
		l.root = l.root.next
		l.root.prev = nil
		node.next = nil
		return
	}

	prev := node.prev

	if node == l.tail {
		prev.next = nil
		l.tail.prev = nil
		l.tail = prev
	} else { // 중간 노드를 지울 때
		node.prev = nil
		prev.next = prev.next.next
		prev.next.prev = prev
	}
	node.next = nil
}

func (l *DoubleLinkedList) PrintNodes() {
	node := l.root
	for node.next != nil {
		fmt.Printf("%d ->", node.val)
		node = node.next
	}
	fmt.Printf("%d\n", node.val)
}

func (l *DoubleLinkedList) PrintReverse() {
	node := l.tail
	for node.prev != nil {
		fmt.Printf("%d ->", node.val)
		node = node.prev
	}
	fmt.Printf("%d\n", node.val)
}

func main() {
	list := &DoubleLinkedList{}
	list.AddNode(0)

	for i := 1; i < 10; i++ {
		list.AddNode(i)
	}

	list.PrintNodes()

	list.RemoveNode(list.root.next)

	list.PrintNodes()

	list.RemoveNode(list.root)

	list.PrintNodes()

	list.RemoveNode(list.tail)

	list.PrintNodes()

	list.PrintReverse()
}
