package ch28stackandqueue

import "golangTraining/2020study/ch27packaging"

type Stack struct {
	l &ch27packaging.*List
}

func (s * Stack) Push(val int){
	s.l.AddNode(val)	
}