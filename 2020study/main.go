package main

import (
	"fmt"
	"golangTraining/2020study/ch27packaging"
)

func main() {
	list := &ch27packaging.DoubleLinkedList{}
	list.AddNode(0)

	for i := 1; i < 10; i++ {
		list.AddNode(i)
	}

	list.PrintNodes()

	list.RemoveNode(list.Root.Next)

	list.PrintNodes()

	list.RemoveNode(list.Root)

	list.PrintNodes()

	list.RemoveNode(list.Tail)

	list.PrintNodes()

	list.PrintReverse()

	stack2 := ch27packaging.NewStack()

	for i := 0; i < 100; i++ {
		stack2.Push(i)
	}

	fmt.Println("New Stack")

	for !stack2.Empty() {
		val := stack2.Pop()
		fmt.Printf("%d ->", val)
	}

}
