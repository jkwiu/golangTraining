package tdd

import (
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// linked list test code
var list *List

func TestNewList(t *testing.T) {
	assert := assert.New(t)
	list = &List{root: &Node{Val: 0}}
	list.tail = list.root
	assert.Equal(list, NewList(0))
}

func TestAddNode(t *testing.T) {
	assert := assert.New(t)
	list.AddNode(1)
	list.PrintNodes()
	assert.Equal(1, list.tail.Val)

	list.AddNode(2)
	list.PrintNodes()
	assert.Equal(2, list.tail.Val)

	list.AddNode(3)
	list.PrintNodes()
	assert.Equal(3, list.tail.Val)
	assert.Equal(0, list.root.Val)
	assert.Equal(1, list.tail.Prev.Prev.Val)
}

func TestRemoveNode(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(0, list.RemoveNode(list.root))
}

// stack test code
var stack *Stack

func TestNewStack(t *testing.T) {
	assert := assert.New(t)
	stack = NewStack(0)
	assert.Equal(0, stack.list.root.Val)
	assert.Equal(0, stack.list.tail.Val)

}

func TestPush_Stack(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(1, stack.Push(1))
	assert.Equal(1, stack.list.tail.Val)
}

func TestPop_Stack(t *testing.T) {
	assert := assert.New(t)
	lastValue := stack.list.tail.Val
	assert.Equal(lastValue, stack.Pop())
}

func TestPrintStack(t *testing.T) {
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.PrintStack()
}

// queue test code
var queue *Queue

func TestNewQueue(t *testing.T) {
	assert := assert.New(t)
	queue = NewQueue(0)
	assert.Equal(0, queue.list.root.Val)
	assert.Equal(0, queue.list.tail.Val)
}

func TestPush_Queue(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(1, queue.Push(1))
	assert.Equal(2, queue.Push(2))
	assert.Equal(3, queue.Push(3))
}

func TestPop_Queue(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(0, queue.Pop())
	assert.Equal(1, queue.Pop())
}

func TestPrintQueue(t *testing.T) {
	queue.Push(4)
	queue.Push(5)
	queue.Push(6)
	stack.PrintStack()
}

// tree test code
//		0
//     /|\
//	 1	2  3
//	 /\ /\ /\
//  4 5 6 7 8 9
var tree *Tree

func TestNewTree(t *testing.T) {
	assert := assert.New(t)
	tree = NewTree()
	assert.Equal(0, tree.Root.Val)
	assert.Equal(3, tree.Root.Childs[2].Val)
}

// Binary Tree
var bt *BinaryTree

func TestAddNode_BinaryTree(t *testing.T) {
	assert := assert.New(t)
	bt = NewBinaryTree(5)
	bt.Root.AddNode(1)
	bt.Root.AddNode(9)
	assert.Equal(9, bt.Root.Right.Val)
	assert.Equal(1, bt.Root.Left.Val)
}

func TestSearch(t *testing.T) {
	assert := assert.New(t)
	binaryTree := NewBinaryTree(5)
	binaryTree.Root.AddNode(3)
	binaryTree.Root.AddNode(2)
	binaryTree.Root.AddNode(4)
	binaryTree.Root.AddNode(8)
	binaryTree.Root.AddNode(7)
	binaryTree.Root.AddNode(6)
	binaryTree.Root.AddNode(10)
	binaryTree.Root.AddNode(9)

	if found, cnt := binaryTree.Search(5); found {
		assert.Equal(true, found)
		assert.Equal(1, cnt)
	} else {
		assert.Equal(false, found)
		assert.Equal(1, cnt)
	}

	if found, cnt := binaryTree.Search(4); found {
		assert.Equal(true, found)
		assert.Equal(3, cnt)
	} else {
		assert.Equal(false, found)
		assert.Equal(3, cnt)
	}

	if found, cnt := binaryTree.Search(11); found {
		assert.Equal(true, found)
		assert.Equal(3, cnt)
	} else {
		assert.Equal(false, found)
		assert.Equal(3, cnt)
	}
}

// Heap
var h *Heap

func TestPush_Heap(t *testing.T) {
	assert := assert.New(t)
	h = &Heap{}
	h.Push(1)
	h.Push(2)
	assert.Equal(2, h.list[0])

	h.Push(3)
	assert.Equal(3, h.list[0])

	h.Push(4)
	assert.Equal(4, h.list[0])
}

func TestPop_Heap(t *testing.T) {
	assert := assert.New(t)

	h.Pop()
	assert.Equal(3, h.list[0])

	h.Pop()
	assert.Equal(2, h.list[0])

	h.Pop()
	assert.Equal(1, h.list[0])
}

// 정수배열과 정수 N이 주어지면, N번째로 작은 배열 원소를 찾으시오
func TestHeapAlgorithm(t *testing.T) {
	assert := assert.New(t)
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	n := 1
	for i := 0; i < 10; i++ {
		h = &Heap{}
		for i := 0; i < len(nums); i++ {
			h.Push(nums[i])
			if h.Count() > n {
				h.Pop()
			}
		}
		assert.Equal(n, h.Pop())
		n++
	}
}

// hash function
func TestHash(t *testing.T) {
	assert := assert.New(t)
	str1 := "what"
	str2 := "what1"
	str3 := "what2"
	str4 := "what3"
	assert.NotEqual(Hash(str1), Hash(str2), Hash(str3), Hash(str4))
}

func TestMap(t *testing.T) {
	assert := assert.New(t)
	m := NewMap()
	m.Add("hell", "o")
	m.Add("h", "i")
	m.Add("j", "k")
	m.Add("awe", "some")
	assert.Equal("o", m.Get("hell"))
	assert.Equal("i", m.Get("h"))
	assert.Equal("k", m.Get("j"))
	assert.Equal("some", m.Get("awe"))
}

// thread1(account)
func TestThread1_Account(t *testing.T) {
	assert := assert.New(t)
	for i := 0; i < 20; i++ {
		account = append(account, &Account{1000, &sync.Mutex{}})
	}
	globalLock = &sync.Mutex{}
	assert.Equal(20000, GetTotalBalance())
	for i := 0; i < 10; i++ {
		go GoTransfer()
	}

	cnt := 0
	for {
		assert.Equal(20000, GetTotalBalance())
		cnt++
		if cnt == 100 {
			break
		}
	}
}

// thread2(channel)
func TestThread2_Channel(t *testing.T) {
	assert := assert.New(t)

	carChan1 := make(chan Car)
	carChan2 := make(chan Car)
	carChan3 := make(chan Car)

	planeChan1 := make(chan Plane)
	planeChan2 := make(chan Plane)
	planeChan3 := make(chan Plane)

	go StartCarWork(carChan1)
	go StartPlaneWork(planeChan1)
	go MakeTire(carChan1, planeChan1, carChan2, planeChan2)
	go MakeEngine(carChan2, planeChan2, carChan3, planeChan3)

	carCnt := 0
	planeCnt := 0
	for {
		select {
		case result := <-carChan3:
			assert.Equal("Car: "+strconv.Itoa(carCnt)+"Tire_C, Engine_C, ", result.val)
			carCnt++
		case result := <-planeChan3:
			assert.Equal("Plane: "+strconv.Itoa(planeCnt)+"Tire_P, Engine_P, ", result.val)
			planeCnt++
		}
		if carCnt == 10 || planeCnt == 10 {
			break
		}
	}
}

// oop1
func TestOOP1(t *testing.T) {
	assert := assert.New(t)
	bread1 := &Bread{}
	appleJam := &AppleJam{}
	bread1.PutJam(appleJam)
	assert.Equal("bread+Apple", bread1.String())

	bread2 := &Bread{}
	orangeJam := &OrangeJam{}
	bread2.PutJam(orangeJam)
	assert.Equal("bread+Orange", bread2.String())
}

// oop2
func TestOOP2(t *testing.T) {
	assert := assert.New(t)
	var c InterfaceA
	c = &StructA{}
	assert.Equal(4, c.AAA(2))
	assert.Equal("2yo", c.BBB(2))
}
