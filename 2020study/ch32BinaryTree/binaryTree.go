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

func Print(n *Node) {
	fmt.Printf("%d -> ", n.Val)
	if n.left != nil {
		Print(n.left)
	}
	if n.right != nil {
		Print(n.right)
	}
}
