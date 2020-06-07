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

// tree
type TreeNode struct {
	Val    int
	Childs []*TreeNode
}

type Tree struct {
	Root *TreeNode
}

// 자식이 2개인 트리노드를 만들어보자.
func NewTree() *Tree {
	val := 0
	tree := &Tree{Root: &TreeNode{Val: val}}
	fmt.Println(tree.Root.Val)
	val++
	for i := 0; i < 3; i++ {
		tree.Root.Childs = append(tree.Root.Childs, &TreeNode{Val: val})
		val++
	}
	for i := 0; i < len(tree.Root.Childs); i++ {
		for j := 0; j < 2; j++ {
			tree.Root.Childs[i].Childs = append(tree.Root.Childs[i].Childs, &TreeNode{Val: val})
			val++
		}
	}
	return tree
}

// DFS1(재귀)
func (t *Tree) DFS1() {
	DFS1(t.Root)
}

func DFS1(n *TreeNode) {
	fmt.Printf("%d -> ", n.Val)
	for i := 0; i < len(n.Childs); i++ {
		DFS1(n.Childs[i])
	}
}

// DFS2(stack)
func (t *Tree) DFS2() {
	var s []*TreeNode
	s = append(s, t.Root)

	for len(s) > 0 {
		var last *TreeNode
		last, s = s[len(s)-1], s[:len(s)-1]

		fmt.Printf("%d -> ", last.Val)

		for i := len(last.Childs) - 1; i >= 0; i-- {
			s = append(s, last.Childs[i])
		}
	}
}

// BFS(queue)
func (t *Tree) BFS() {
	var q []*TreeNode
	q = append(q, t.Root)

	for len(q) > 0 {
		var first *TreeNode

		first, q = q[0], q[1:]

		fmt.Printf("%d -> ", first.Val)

		for i := 0; i < len(first.Childs); i++ {
			q = append(q, first.Childs[i])
		}
	}
}

// 이진트리
type BinaryTreeNode struct {
	Val   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

type BinaryTree struct {
	Root *BinaryTreeNode
}

type depthNode struct {
	depth int
	node  *BinaryTreeNode
}

// new
func NewBinaryTree(val int) *BinaryTree {
	return &BinaryTree{Root: &BinaryTreeNode{Val: val}}
}

// add
func (b *BinaryTreeNode) AddNode(val int) int {
	if val > b.Val {
		if b.Right != nil {
			b.Right.AddNode(val)
		} else {
			b.Right = &BinaryTreeNode{Val: val}
		}
		return b.Right.Val
	} else {
		if b.Left != nil {
			b.Left.AddNode(val)
		} else {
			b.Left = &BinaryTreeNode{Val: val}
		}
		return b.Left.Val
	}
}

// print
// BFS로 하겠슴돠.
func (b *BinaryTree) Print() {
	q := []depthNode{}
	q = append(q, depthNode{depth: 0, node: b.Root})
	currentDepth := 0

	for len(q) > 0 {
		var first depthNode

		first, q = q[0], q[1:]

		if first.depth != currentDepth {
			fmt.Println()
			currentDepth = first.depth
		}

		fmt.Print(first.node.Val, " ")

		if first.node.Left != nil {
			q = append(q, depthNode{depth: currentDepth + 1, node: first.node.Left})
		}

		if first.node.Right != nil {
			q = append(q, depthNode{depth: currentDepth + 1, node: first.node.Right})
		}
	}
}

// search
func (b *BinaryTree) Search(val int) (bool, int) {
	return b.Root.NodeSearch(val, 1)
}

func (n *BinaryTreeNode) NodeSearch(val int, cnt int) (bool, int) {
	if val == n.Val {
		return true, cnt
	} else if val > n.Val {
		if n.Right != nil {
			return n.Right.NodeSearch(val, cnt+1)
		}
		return false, cnt
	} else {
		if n.Left != nil {
			return n.Left.NodeSearch(val, cnt+1)
		}
		return false, cnt
	}
}

// 최대 heap
type Heap struct {
	list []int
}

// push
func (h *Heap) Push(val int) int {
	h.list = append(h.list, val)
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
	return h.list[0]
}

// pop
func (h *Heap) Pop() int {
	top := h.list[0]
	last := h.list[len(h.list)-1]

	if len(h.list) != 1 {
		h.list = h.list[:len(h.list)-1]
		h.list[0] = last
	} else {
		h.list = h.list[:len(h.list)-1]
	}

	idx := 0
	for idx < len(h.list) {
		swapIdx := -1
		leftIdx := idx*2 + 1

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

		if swapIdx < 0 {
			break
		}
		h.list[idx], h.list[swapIdx] = h.list[swapIdx], h.list[idx]
		idx = swapIdx
	}
	return top
}

// get heap length
func (h *Heap) Count() int {
	return len(h.list)
}

// print heap
func (h *Heap) Print() {
	fmt.Println(h.list)
}
