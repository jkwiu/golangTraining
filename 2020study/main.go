package main

import (
	"fmt"
	"golangTraining/2020study/ch32BinaryTree"
)

func main() {
	bt := ch32BinaryTree.NewBinaryTree(3)

	fmt.Println(bt.Root.Val)

	bt.Root.AddNode(1)
	bt.Root.AddNode(5)
	bt.Root.AddNode(2)
	bt.Root.AddNode(4)
	bt.Root.AddNode(6)

	fmt.Println("DFS")
	bt.PrintDFS(bt.Root)

	fmt.Println()

	fmt.Println("BFS")
	bt.PrintBFS()

	fmt.Println()
	if found, cnt := bt.Search(6); found {
		fmt.Println("found 6 cnt:", cnt)
	} else {
		fmt.Println("not found 6 cnt:", cnt)
	}

	if found, cnt := bt.Search(11); found {
		fmt.Println("found 11 cnt:", cnt)
	} else {
		fmt.Println("not found 11 cnt:", cnt)
	}
}
