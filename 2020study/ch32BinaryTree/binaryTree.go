package ch32BinaryTree

import "fmt"

type BinaryTreeNode struct {
	Val   int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

type BinaryTree struct {
	Root *BinaryTreeNode
}

func NewBinaryTree(v int) *BinaryTree {
	tree := &BinaryTree{}
	tree.Root = &BinaryTreeNode{Val: v}
	return tree
}

func (b *BinaryTreeNode) AddNode(v int) *BinaryTreeNode {
	if b.Val > v {
		if b.left == nil {
			b.left = &BinaryTreeNode{Val: v}
			return b.left
		} else {
			return b.left.AddNode(v)
		}
	} else {
		if b.right == nil {
			b.right = &BinaryTreeNode{Val: v}
			return b.right
		} else {
			return b.right.AddNode(v)
		}
	}
}

type depthNode struct {
	depth int
	node  *BinaryTreeNode
}

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

		if first.node.left != nil {
			q = append(q, depthNode{depth: currentDepth + 1, node: first.node.left})
		}

		if first.node.right != nil {
			q = append(q, depthNode{depth: currentDepth + 1, node: first.node.right})
		}
	}
}

func (b *BinaryTree) Search(v int) (bool, int) {
	return b.Root.Search(v, 1)
}

func (b *BinaryTreeNode) Search(v int, cnt int) (bool, int) {
	if b.Val == v {
		return true, cnt
	} else if b.Val > v {
		if b.left != nil {
			return b.left.Search(v, cnt+1)
		}
		return false, cnt
	} else {
		if b.right != nil {
			return b.right.Search(v, cnt+1)
		}
		return false, cnt
	}
}
