package main

import (
	"fmt"
	"golangTraining/2020study/ch29tree"
)

func main() {
	tree := &ch29tree.Tree{}

	val := 1
	tree.AddNode(val)
	val++
	for i := 0; i < 3; i++ {
		tree.Root.AddChild(val)
		val++
	}

	for i := 0; i < len(tree.Root.Child); i++ {
		for j := 0; j < 2; j++ {
			tree.Root.Child[i].AddChild(val)
			val++
		}
	}

	tree.DFS1()

	fmt.Println()

	tree.DFS2()

	tree.BFS()
}
