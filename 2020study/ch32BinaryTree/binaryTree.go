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

// 여기서 반환값을 사용하지 않음에도 지정해줌에 주목할 필요가 있다.
// 그건 바로 함수의 완결성을 위해서이다.
// 호출하는 쪽에서 이 값을 사용할지 말지를 정할 수 있는 옵션을 주는 것이다.
// 함수를 짤 때는 한가지의 목적보다는 여러가지 가능성을 염두해 두고 짜는 것이 좋다.
// from goSense
func (n *Node) AddNode(val int) *Node {
	if val < n.Val {
		if n.left == nil {
			n.left = &Node{Val: val}
			return n.left
		} else {
			return n.left.AddNode(val)
		}
	} else {
		if n.right == nil {
			n.right = &Node{Val: val}
			return n.right
		} else {
			return n.right.AddNode(val)
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
