/*
//방법 1`
package main

import "fmt"

//맨 처음 노드를 알고 있어야 다음 노드로 넘어갈 수 있다.
//맨 처음 노드는 root

type Node struct {
  next *Node
  val int
}



func main()  {
  var root *Node

  root = &Node{val:0}

  for i:=1; i<10; i++{
    AddNode(root, i)
  }

  node := root
  for node.next != nil{
    fmt.Printf("%d ->", node.val)
    node = node.next
  }
  fmt.Printf("%d\n", node.val)


}

func AddNode(root *Node, val int)  {
  var tail *Node
  tail = root
  //tail의 next가 존재하는 이상 다음 next를 가르킨다. next==nil일 때까지
  for tail.next != nil {
    tail = tail.next
  }

  node := &Node{val:val}
  tail.next = node
}
*/

//방법2
package main

import "fmt"

type Node struct {
	next *Node
	val  int
}

func main() {
	var root *Node
	var tail *Node

	root = &Node{val: 0}
	tail = root

	for i := 1; i < 10; i++ {
		tail = AddNode(tail, i)
	}

	PrintNodes(root)

	root, tail = RemoveNode(root.next, root, tail)

	PrintNodes(root)

	root, tail = RemoveNode(root, root, tail)

	PrintNodes(root)

	root, tail = RemoveNode(tail, root, tail)

	PrintNodes(root)
	fmt.Printf("tail:%d\n", tail.val)
}

func AddNode(tail *Node, val int) *Node {
	node := &Node{val: val}
	tail.next = node
	return node
}

//Node를 지우고자 할 때
func RemoveNode(node *Node, root *Node, tail *Node) (*Node, *Node) {

	if node == root {
		root = root.next
		if root == nil {
			tail = nil
		}
		return root, tail
	}

	prev := root
	for prev.next != node {
		prev = prev.next
	}
	if node == tail {
		prev.next = nil
		tail = prev
	} else {
		prev.next = prev.next.next
	}
	return root, tail
}

func PrintNodes(root *Node) {
	node := root
	for node.next != nil {
		fmt.Printf("%d ->", node.val)
		node = node.next
	}
	fmt.Printf("%d\n", node.val)
}
