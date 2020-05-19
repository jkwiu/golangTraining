package ch32BinaryTree

import "fmt"

type Node struct {
	Val   int
	left  *Node
	right *Node
}

type BinaryTree struct {
	Root *Node
}

func NewBinaryTree(val int) *BinaryTree {
	return &BinaryTree{Root: &Node{Val: val}}
}

func (n *Node) AddNode(val int) {
	if val < n.Val {
		if n.left == nil {
			n.left = &Node{Val: val}
		} else {
			n.left.AddNode(val)
		}
	} else {
		if n.right == nil {
			n.right = &Node{Val: val}
		} else {
			n.right.AddNode(val)
		}
	}
}

// DFS
// 재귀함수로 구현
func (b *BinaryTree) PrintDFS(n *Node) {
	fmt.Printf("%d -> ", n.Val)
	if n.left != nil {
		b.PrintDFS(n.left)
	}
	if n.right != nil {
		b.PrintDFS(n.right)
	}
}

type depthNode struct {
	depth int
	node  *Node
}

// BFS
func (b *BinaryTree) PrintBFS() {
	// queue로 구현
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

		if first.node.left != nil {
			q = append(q, depthNode{depth: currentDepth + 1, node: first.node.left})
		}
		if first.node.right != nil {
			q = append(q, depthNode{depth: currentDepth + 1, node: first.node.right})
		}
	}
}

func (b *BinaryTree) Search(val int) (bool, int) {
	return b.Root.NodeSearch(val, 1)
}

func (n *Node) NodeSearch(val int, cnt int) (bool, int) {
	if val == n.Val {
		return true, cnt
	} else if val > n.Val {
		if n.right != nil {
			return n.right.NodeSearch(val, cnt+1)
		}
		return false, cnt
	} else {
		if n.left != nil {
			return n.left.NodeSearch(val, cnt+1)
		}
		return false, cnt
	}
}
