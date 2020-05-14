package ch29.tree

type TreeNode struct {
	Val int
	Childs []*TreeNode
}

type Tree struct {
	Root *TreeNode
}

func (t *Tree) AddNode(val int){
	
}