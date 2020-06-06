package tdd

import (
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
