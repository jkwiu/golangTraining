package main

import "fmt"

// 맨 처음 노드를 항상 알고 있어야 한다.
// 두가지 방법이 있는데 tail을 알고 있으면 처음부터 끝까지 진행하여 tail을 찾을 필요가 없다.
type Node struct {
	next *Node
	val  int
}

// OOP / 응징섭은 높이고 종속성은 낮추는 작업
type LinkedList struct {
	root *Node
	tail *Node
}

// method 시작
func (l *LinkedList) AddNode(val int) {
	// node가 없을 때 하나를 생성
	if l.root == nil {
		l.root = &Node{val: val}
		l.tail = l.root
		return
	}
	l.tail.next = &Node{val: val}
	l.tail = l.tail.next
}

func (l *LinkedList) RemoveNode(node *Node) {
	if node == l.root {
		l.root = l.root.next
		node.next = nil
		return
	}

	prev := l.root
	for prev.next != node {
		prev = prev.next
	}

	if node == l.tail {
		prev.next = nil
		l.tail = prev
	} else { // 중간 노드를 지울 때
		prev.next = prev.next.next
	}
	node.next = nil
}

func (l *LinkedList) PrintNodes() {
	node := l.root
	for node.next != nil {
		fmt.Printf("%d ->", node.val)
		node = node.next
	}
	fmt.Printf("%d\n", node.val)
}

// method 끝

func main() {
	list := &LinkedList{}
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

	fmt.Println("/////////////////////////")

	var root *Node
	var tail *Node

	root = &Node{val: 0}
	tail = root

	for i := 1; i < 10; i++ {
		tail = AddNode(tail, i)
	}

	PrintNodes(root)

	root, tail = RemoveNode(root.next, root, tail)

	PrintNodes(root)

	root, tail = RemoveNode(root, root, tail)

	PrintNodes(root)

	root, tail = RemoveNode(tail, root, tail)

	PrintNodes(root)
	fmt.Printf("tail:%d\n", tail.val)

}

// 전자의 방법대로 해서 add를 하고자 하면 n개이면 n번 돌아야 한다. O(N)
// 후자의 방법은 O(1)
func AddNode(tail *Node, val int) *Node {
	node := &Node{val: val}
	tail.next = node
	return node
}

// 지우고자 하는 노드를 가르키는 것을 건너뛰면 삭제됨(gc가 삭제한다)
// 맨 처음 노드를 지우고자 할 때는 root만 바꿔주면 된다 O(1)
// 중간 노드를 지울 때 O(N)
// 맨 끝 노드를 지우고자 할 때는 tail만 바꿔주면 된다 O(1)
func RemoveNode(node *Node, root *Node, tail *Node) (*Node, *Node) {

	// 맨 처음 노드를 지우고자 할 때
	if node == root {
		root = root.next
		if root == nil {
			tail = nil
		}
		return root, tail
	}

	// 이전 노드를 찾는다 O(N) 100만개면 100만번을 돌려야 한다....그래서 double linked list로 가보자
	prev := root
	for prev.next != node {
		prev = prev.next
	}

	// 맨 끝 노드를 지우고자 할 때
	if node == tail {
		prev.next = nil
		tail = prev
	} else { // 중간 노드를 지울 때
		prev.next = prev.next.next
	}

	return root, tail
}

// 노드들을 출력
func PrintNodes(root *Node) {
	node := root
	for node.next != nil {
		fmt.Printf("%d ->", node.val)
		node = node.next
	}
	fmt.Printf("%d\n", node.val)
}
