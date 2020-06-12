// double linked list
// stack
// queue
// tree
// DFS1(재귀), 깊이 우선 탐색
// DFS2(스택)
// BFS(큐), 너비 우선 탐색
// binary tree
// BST(Binary Search Tree)
// 최소힙
// 힙을 이용한 알고리즘
// hash
// map

package tdd

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// 1. linked list
type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

type List struct {
	root *Node
	tail *Node
}

func NewList(val int) *List {
	list := &List{root: &Node{Val: val}}
	list.tail = list.root
	return list
}

func (l *List) AddNode(val int) int {
	if l.root == nil {
		fmt.Println("There is no List!!")
		return 0
	}
	l.tail.Next = &Node{Val: val}
	l.tail.Next.Prev = l.tail
	l.tail = l.tail.Next
	return l.tail.Val
}

func (l *List) RemoveNode(n *Node) int {
	if l.root == nil {
		fmt.Println("There is no List!!")
		return 0
	}
	rst := n.Val
	if n == l.root { // 맨 앞일 때
		l.root = l.root.Next
		l.root.Prev.Next = nil
		l.root.Prev = nil
	}
	return rst
}

func (l *List) PrintNodes() {
	node := l.root
	for node.Next != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}
	fmt.Println(node.Val)
}

// stack
type Stack struct {
	list *List
}

// 신규
func NewStack(val int) *Stack {
	stack := &Stack{list: &List{root: &Node{Val: val}}}
	stack.list.tail = stack.list.root
	return stack
}

// 추가
func (s *Stack) Push(val int) int {
	return s.list.AddNode(val)
}

// 삭제
func (s *Stack) Pop() int {
	return s.list.RemoveNode(s.list.tail)
}

// 출력
func (s *Stack) PrintStack() {
	node := s.list.tail
	for node.Prev != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Prev
	}
	fmt.Println(node.Val)
}

// queue
type Queue struct {
	list *List
}

// 신규
func NewQueue(val int) *Queue {
	queue := &Queue{list: &List{root: &Node{Val: val}}}
	queue.list.tail = queue.list.root
	return queue
}

// 추가
func (q *Queue) Push(val int) int {
	return q.list.AddNode(val)
}

// 삭제
func (q *Queue) Pop() int {
	return q.list.RemoveNode(q.list.root)
}

// 출력
func (q *Queue) PrintQueue() {
	q.list.PrintNodes()
}

// tree
type TreeNode struct {
	Val    int
	Childs []*TreeNode
}

type Tree struct {
	Root *TreeNode
}

// 자식이 2개인 트리노드를 만들어보자.
func NewTree() *Tree {
	val := 0
	tree := &Tree{Root: &TreeNode{Val: val}}
	fmt.Println(tree.Root.Val)
	val++
	for i := 0; i < 3; i++ {
		tree.Root.Childs = append(tree.Root.Childs, &TreeNode{Val: val})
		val++
	}
	for i := 0; i < len(tree.Root.Childs); i++ {
		for j := 0; j < 2; j++ {
			tree.Root.Childs[i].Childs = append(tree.Root.Childs[i].Childs, &TreeNode{Val: val})
			val++
		}
	}
	return tree
}

// DFS1(재귀)
func (t *Tree) DFS1() {
	DFS1(t.Root)
}

func DFS1(n *TreeNode) {
	fmt.Printf("%d -> ", n.Val)
	for i := 0; i < len(n.Childs); i++ {
		DFS1(n.Childs[i])
	}
}

// DFS2(stack)
func (t *Tree) DFS2() {
	var s []*TreeNode
	s = append(s, t.Root)

	for len(s) > 0 {
		var last *TreeNode
		last, s = s[len(s)-1], s[:len(s)-1]

		fmt.Printf("%d -> ", last.Val)

		for i := len(last.Childs) - 1; i >= 0; i-- {
			s = append(s, last.Childs[i])
		}
	}
}

// BFS(queue)
func (t *Tree) BFS() {
	var q []*TreeNode
	q = append(q, t.Root)

	for len(q) > 0 {
		var first *TreeNode

		first, q = q[0], q[1:]

		fmt.Printf("%d -> ", first.Val)

		for i := 0; i < len(first.Childs); i++ {
			q = append(q, first.Childs[i])
		}
	}
}

// 이진트리
type BinaryTreeNode struct {
	Val   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

type BinaryTree struct {
	Root *BinaryTreeNode
}

type depthNode struct {
	depth int
	node  *BinaryTreeNode
}

// new
func NewBinaryTree(val int) *BinaryTree {
	return &BinaryTree{Root: &BinaryTreeNode{Val: val}}
}

// add
func (b *BinaryTreeNode) AddNode(val int) int {
	if val > b.Val {
		if b.Right != nil {
			b.Right.AddNode(val)
		} else {
			b.Right = &BinaryTreeNode{Val: val}
		}
		return b.Right.Val
	} else {
		if b.Left != nil {
			b.Left.AddNode(val)
		} else {
			b.Left = &BinaryTreeNode{Val: val}
		}
		return b.Left.Val
	}
}

// print
// BFS로 하겠슴돠.
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

		if first.node.Left != nil {
			q = append(q, depthNode{depth: currentDepth + 1, node: first.node.Left})
		}

		if first.node.Right != nil {
			q = append(q, depthNode{depth: currentDepth + 1, node: first.node.Right})
		}
	}
}

// search
func (b *BinaryTree) Search(val int) (bool, int) {
	return b.Root.NodeSearch(val, 1)
}

func (n *BinaryTreeNode) NodeSearch(val int, cnt int) (bool, int) {
	if val == n.Val {
		return true, cnt
	} else if val > n.Val {
		if n.Right != nil {
			return n.Right.NodeSearch(val, cnt+1)
		}
		return false, cnt
	} else {
		if n.Left != nil {
			return n.Left.NodeSearch(val, cnt+1)
		}
		return false, cnt
	}
}

// 최대 heap
type Heap struct {
	list []int
}

// push
func (h *Heap) Push(val int) int {
	h.list = append(h.list, val)
	idx := len(h.list) - 1
	for idx >= 0 {
		parentIdx := (idx - 1) / 2
		if parentIdx < 0 {
			break
		}
		if h.list[idx] > h.list[parentIdx] {
			h.list[idx], h.list[parentIdx] = h.list[parentIdx], h.list[idx]
			idx = parentIdx
		} else {
			break
		}
	}
	return h.list[0]
}

// pop
func (h *Heap) Pop() int {
	top := h.list[0]
	last := h.list[len(h.list)-1]

	if len(h.list) != 1 {
		h.list = h.list[:len(h.list)-1]
		h.list[0] = last
	} else {
		h.list = h.list[:len(h.list)-1]
	}

	idx := 0
	for idx < len(h.list) {
		swapIdx := -1
		leftIdx := idx*2 + 1

		if leftIdx >= len(h.list) {
			break
		}

		if h.list[leftIdx] > h.list[idx] {
			swapIdx = leftIdx
		}

		rightIdx := idx*2 + 2
		if rightIdx < len(h.list) {
			if h.list[rightIdx] > h.list[idx] {
				if swapIdx < 0 || h.list[swapIdx] < h.list[rightIdx] {
					swapIdx = rightIdx
				}
			}
		}

		if swapIdx < 0 {
			break
		}
		h.list[idx], h.list[swapIdx] = h.list[swapIdx], h.list[idx]
		idx = swapIdx
	}
	return top
}

// get heap length
func (h *Heap) Count() int {
	return len(h.list)
}

// print heap
func (h *Heap) Print() {
	fmt.Println(h.list)
}

// map
func Hash(s string) int {
	h := 0
	A := 256
	B := 3571
	for i := 0; i < len(s); i++ {
		h = (h*A + int(s[i])) % B
	}
	return h
}

type keyValue struct {
	key   string
	value string
}

type Map struct {
	keyArray [3571][]keyValue
}

func NewMap() *Map {
	return &Map{}
}

func (m *Map) Add(key, value string) {
	h := Hash(key)
	m.keyArray[h] = append(m.keyArray[h], keyValue{key, value})
}

func (m *Map) Get(key string) string {
	h := Hash(key)
	for i := 0; i < len(m.keyArray[h]); i++ {
		if m.keyArray[h][i].key == key {
			return m.keyArray[h][i].value
		}
	}
	return ""
}

// thread1(mutex)
type Account struct {
	balance int
	mutex   *sync.Mutex
}

var account []*Account
var globalLock *sync.Mutex

func (a *Account) Withdraw(val int) {
	a.mutex.Lock()
	a.balance -= val
	a.mutex.Unlock()
}

func (a *Account) Deposit(val int) {
	a.mutex.Lock()
	a.balance += val
	a.mutex.Unlock()
}

func (a *Account) Balance() int {
	a.mutex.Lock()
	balance := a.balance
	a.mutex.Unlock()
	return balance
}

func Transfer(sender, receiver int, money int) {
	globalLock.Lock()
	account[sender].Withdraw(money)
	account[receiver].Deposit(money)
	globalLock.Unlock()
}

func RandomTransfer() {
	var sender, balance int
	for {
		sender = rand.Intn(len(account))
		balance = account[sender].Balance()
		if balance > 0 {
			break
		}
	}

	var receiver int
	for {
		receiver = rand.Intn(len(account))
		if sender != receiver {
			break
		}
	}
	money := rand.Intn(balance)
	Transfer(sender, receiver, money)
}

func GoTransfer() {
	for {
		RandomTransfer()
	}
}

func GetTotalBalance() int {
	globalLock.Lock()
	total := 0
	for i := 0; i < len(account); i++ {
		total += account[i].balance
	}
	globalLock.Unlock()
	return total
}

func PrintTotalBalance() {
	fmt.Printf("Total: %d\n", GetTotalBalance())
}

// thread2(channel)
// 컨베이 벨트 방식
type Car struct {
	val string
}

type Plane struct {
	val string
}

func StartCarWork(chan1 chan Car) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		chan1 <- Car{val: "Car: " + strconv.Itoa(i)}
		i++
	}
}
func StartPlaneWork(chan1 chan Plane) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		chan1 <- Plane{val: "Plane: " + strconv.Itoa(i)}
		i++
	}
}

func MakeTire(carChan chan Car, planeChan chan Plane, outCarChan chan Car, outPlaneChan chan Plane) {
	for {
		select {
		case car := <-carChan:
			car.val += "Tire_C, "
			outCarChan <- car
		case plane := <-planeChan:
			plane.val += "Tire_P, "
			outPlaneChan <- plane
		}
	}
}

func MakeEngine(carChan chan Car, planeChan chan Plane, outCarChan chan Car, outPlaneChan chan Plane) {
	for {
		select {
		case car := <-carChan:
			car.val += "Engine_C, "
			outCarChan <- car
		case plane := <-planeChan:
			plane.val += "Engine_P, "
			outPlaneChan <- plane
		}
	}
}

// OOP1
type Bread struct {
	val string
}

type Jam interface {
	GetOneSpoon() SpoonOfJam
}

type SpoonOfJam interface {
	String() string
}

func (b *Bread) PutJam(jam Jam) {
	spoon := jam.GetOneSpoon()
	b.val += spoon.String()
}

func (b *Bread) String() string {
	return "bread" + b.val
}

type AppleJam struct {
}

type SpoonOfAppleJam struct {
}

func (a *AppleJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfAppleJam{}
}

func (a *SpoonOfAppleJam) String() string {
	return "+Apple"
}

type OrangeJam struct {
}

type SpoonOfOrangeJam struct {
}

func (a *OrangeJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfOrangeJam{}
}

func (a *SpoonOfOrangeJam) String() string {
	return "+Orange"
}

// oop2
type InterfaceA interface {
	AAA(int) int
	BBB(int) string
}

type StructA struct {
}

func (s *StructA) AAA(x int) int {
	return x * x
}

func (s *StructA) BBB(x int) string {
	return strconv.Itoa(x) + "yo"
}
