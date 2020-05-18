package ch29tree

import "fmt"

type Node struct {
	Val   int
	Child []*Node
}

type Tree struct {
	Root *Node
}

func (t *Tree) AddNode(val int) {
	if t.Root == nil {
		t.Root = &Node{Val: val}
	} else {
		t.Root.Child = append(t.Root.Child, &Node{Val: val})
	}
}

func (n *Node) AddChild(val int) {
	n.Child = append(n.Child, &Node{Val: val})
}

// DFS
// 1. 재귀호출
// 자식들을 돌면서 그 자식들을 길이만큼 호출한다
func (t *Tree) DFS1() {
	DFS1(t.Root)
}

func DFS1(n *Node) {
	fmt.Printf("%d -> ", n.Val)
	for i := 0; i < len(n.Child); i++ {
		DFS1(n.Child[i])
	}
}

// 2. stack
func (t *Tree) DFS2() {
	s := []*Node{}
	s = append(s, t.Root)

	for len(s) > 0 {
		var last *Node
		last, s = s[len(s)-1], s[:len(s)-1]

		fmt.Printf("%d -> ", last.Val)

		for i := len(last.Child) - 1; i >= 0; i-- {
			s = append(s, last.Child[i])
		}
	}
	fmt.Println()
}

// BFS
// 1. QUEUE
func (t *Tree) BFS() {
	queue := []*Node{}
	queue = append(queue, t.Root)

	for len(queue) > 0 {
		var first *Node
		first, queue = queue[0], queue[1:]

		fmt.Printf("%d -> ", first.Val)

		for i := 0; i < len(first.Child); i++ {
			queue = append(queue, first.Child[i])
		}
	}
}
