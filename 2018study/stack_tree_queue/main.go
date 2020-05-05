package main

import (
	"fmt"
	"training/golang_lecture/stack_tree_queue/dataStruct"
)

func main() {
	tree := dataStruct.Tree{}

	val := 1
	tree.AddNode(val)
	val++

	for i := 0; i < 3; i++ {
		tree.Root.AddNode(val)
		val++
	}

	for i := 0; i < len(tree.Root.Childs); i++ {
		for j := 0; j < 2; j++ {
			tree.Root.Childs[i].AddNode(val)
			val++
		}
	}

	fmt.Println("# DFS, BFS, BTS 수업")
	tree.DFS1()
	fmt.Println()

	tree.DFS2()
	fmt.Println()

	tree.BFS()
	fmt.Println()

	tree2 := dataStruct.NewBinaryTree(5)
	tree2.Root.AddNode(3)
	tree2.Root.AddNode(2)
	tree2.Root.AddNode(4)
	tree2.Root.AddNode(8)
	tree2.Root.AddNode(7)
	tree2.Root.AddNode(6)
	tree2.Root.AddNode(10)
	tree2.Root.AddNode(9)

	tree2.Print()
	fmt.Println()

	if found, cnt := tree2.Search(6); found {
		fmt.Println("found 6 cnt:", cnt)
	} else {
		fmt.Println("not found 6 cnt:", cnt)
	}

	if found, cnt := tree2.Search(11); found {
		fmt.Println("found 11 cnt:", cnt)
	} else {
		fmt.Println("not found 11 cnt:", cnt)
	}

	fmt.Println("# Heap 수업")
	//Heap 항목 추가
	h := &dataStruct.Heap{}

	h.Push(2)
	h.Push(6)
	h.Push(9)
	h.Push(6)
	h.Push(7)
	h.Push(8)

	h.Print()

	//Heap 정렬
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())

	fmt.Println("# Map 수업")
	//map.go
	fmt.Println("abced = ", dataStruct.Hash("abcde"))
	fmt.Println("abced = ", dataStruct.Hash("abcde"))
	fmt.Println("abcef = ", dataStruct.Hash("abcdf"))
	fmt.Println("tbced = ", dataStruct.Hash("tbcde"))
	fmt.Println("abcedfdfdfd = ", dataStruct.Hash("abcdefdfdfd"))

	m := dataStruct.CreatMap()
	m.Add("AAA", "0107777777")
	m.Add("BBB", "0108888888")
	m.Add("CCC", "0109999999")
	m.Add("ASDFSA", "0101111111")

	fmt.Println("AAA = ", m.Get("AAA"))
	fmt.Println("BBB = ", m.Get("BBB"))
	fmt.Println("CCC = ", m.Get("CCC"))
	fmt.Println("DDD = ", m.Get("DDD"))
	fmt.Println("ASDFSA = ", m.Get("ASDFSA"))
}
