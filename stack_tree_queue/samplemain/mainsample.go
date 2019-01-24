package main

import "fmt"

type Node struct {
	next *Node
	prev *Node
	val int
}

type LinkedList struct {
	root *Node
	tail *Node
}

func (l *LinkedList) AddNode(val int) {
	if l.root == nil {
		l.root = &Node{val:val}
		l.tail = l.root
		return
	}
	prev := l.tail
	l.tail.next = &Node{val:val}
	l.tail = l.tail.next
	l.tail.prev = prev
}

func (l *LinkedList) RemoveNode(node *Node) {
	if node == l.root {
		l.root = l.root.next
		if l.root != nil{
		l.root.prev = nil
	}
		node.next=nil
		return
	}

	prev := node.prev
	if node == l.tail{
		l.tail.prev=nil
		prev.next=nil
		l.tail=prev
	} else {
		node.prev=nil
		prev.next=prev.next.next
		prev.next=prev
	}
	node.next=nil
}

func (l *LinkedList) PrintNodes()  {
	node := l.root
	for node.next != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.next
	}
	fmt.Printf("%d\n", node.val)
}

func (l *LinkedList) PrintReverse() {
	node := l.tail
	for node.prev != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.prev
	}
	fmt.Printf("%d\n", node.val)
}

func (l *LinkedList) Back() int {
	if l.tail == nil {
		return 0
	}
	return l.tail.val
}

func (l *LinkedList) Front() int {
	if l.root == nil {
		return 0
	}
	return l.root.val
}

func (l *LinkedList) PopBack() {
	if l.tail == nil {
		return
	}
	l.RemoveNode(l.tail)
}

func (l *LinkedList) PopFront() {
	if l.root == nil {
		return
	}
	l.RemoveNode(l.root)
}

func (l *LinkedList) Empty() bool {
	return l.root == nil
}

//--------------queue---------------------
type Queue struct {
	ll *LinkedList
}

func NewQueue() *Queue {
	return &Queue{ll:&LinkedList{}}
}

func (q *Queue) Push(val int) {
	q.ll.AddNode(val)
}

func (q *Queue) Pop() int {
	front := q.ll.Front()
	q.ll.PopFront()
	return front
}

func (q *Queue) Empty() bool {
	return q.ll.Empty()
}

//-----------Stack-------------------
type Stack struct {
	ll *LinkedList
}

func NewStack() *Stack {
	return &Stack{ll:  &LinkedList{}}
}

func (s *Stack) Empty() bool {
	return s.ll.Empty()
}

func (s *Stack) Push(val int) {
	s.ll.AddNode(val)
}

func (s *Stack) Pop() int {
	back := s.ll.Back()
	s.ll.PopBack()
	return back
}

//--------------TreeNode--------------
type TreeNode struct {
	Val int
	Childs []*TreeNode
}

type Tree struct {
	Root *TreeNode
}

func (t *Tree) AddNode(val int) {
	if t.Root == nil {
		t.Root = &TreeNode{Val:val}
	} else {
		t.Root.Childs = append(t.Root.Childs, &TreeNode{Val:val})
	}
}

func (t *TreeNode) AddNode(val int) {
	t.Childs = append(t.Childs, &TreeNode{Val: val})
}

//DFS(Depth Force Search)
// 1. 재귀호출
func DFS1(node *TreeNode)  {
	fmt.Printf("%d->", node.Val)

	for i:=0;i<len(node.Childs); i++{
		DFS1(node.Childs[i])
	}
}

func (t *Tree) DFS1() {
	DFS1(t.Root)
}

//DFS2
//2. Stack을 이용한 방법
func (t *Tree) DFS2() {
	s := []*TreeNode{}
	s = append(s, t.Root)

	for len(s)>0{
		var last *TreeNode
		last, s = s[len(s)-1], s[:len(s)-1]

		fmt.Printf("%d->",last.Val)

		for i:=len(last.Childs)-1; i>=0; i--{
			s = append(s, last.Childs[i])
		}
	}
}

//BFS
//1. Queue
func (t *Tree) BFS() {
	queue := []*TreeNode{}
	queue = append(queue, t.Root)

	for len(queue)>0 {
		var first *TreeNode
		first, queue=queue[0], queue[1:]

		fmt.Printf("%d->",first.Val)
		for i:=0;i<len(first.Childs); i++{
			queue = append(queue, first.Childs[i])
		}
	}
}

//Heap
type Heap struct {
	list []int
}

func (h *Heap) Push(v int) {
	h.list = append(h.list, v)

	idx := len(h.list)-1
	for idx>=0{
		parentIdx := (idx-1)/2
		if parentIdx < 0 {
			break;
		}

	}
}

//------------Main------------------------
func main()  {

	//slice stack add
	stack1 := []int{}
	for i:=0;i<5;i++{
		stack1=append(stack1, i)
	}
	fmt.Println(stack1)

	//slice stack delete
	for len(stack1)>0{
	var last int
	last, stack1 = stack1[len(stack1)-1], stack1[:len(stack1)-1]
	fmt.Println(last)
}

//slice queue add
queue1 := []int{}
for i:=0;i<5;i++{
	queue1=append(queue1, i)
}
fmt.Println(queue1)

//slice queue delete
for len(queue1)>0{
var first int
first, queue1 = queue1[0], queue1[1:]
fmt.Println(first)
}

fmt.Println("stack")
//Linked List stack
stack2 := NewStack()
for i:=0;i<5;i++{
	stack2.Push(i)
	}
	stack2.ll.PrintNodes()
	for !stack2.Empty(){
		val := stack2.Pop()
		fmt.Printf("%d -> ", val)
	}

	fmt.Println()
	fmt.Println("Queue")
	//Linked List queue
	queue2 := NewQueue()
	for i:=0;i<5;i++{
		queue2.Push(i)
		}
		queue2.ll.PrintNodes()
		for !queue2.Empty(){
			val := queue2.Pop()
			fmt.Printf("%d -> ", val)
		}
		fmt.Println()

		//DFS1
		fmt.Println("DFS1 재귀호출")
		tree := &Tree{}
		val:=1
		tree.AddNode(val)
		val++

		for i:=0;i<3;i++{
			tree.Root.AddNode(val)
			val++
		}

		for i:=0;i<len(tree.Root.Childs);i++{
			for j:=0; j<2; j++{
				tree.Root.Childs[i].AddNode(val)
				val++
			}
		}
		tree.DFS1()

		fmt.Println()

		//DFS2
		fmt.Println("DFS2 Stack")
		tree.DFS2()
		fmt.Println()

		//BFS
		fmt.Println("BFS Queue")
		tree.BFS()
		fmt.Println()
}
