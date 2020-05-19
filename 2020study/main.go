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

	ch32BinaryTree.Print(bt.Root)
}
