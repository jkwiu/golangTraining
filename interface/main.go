package main

import (
	"fmt"
	"strconv"
)

type StructA struct {
	val string
}

func (s *StructA) String() string {
	return "Val:" + s.val
}

type StructB struct {
	val int
}

func (s *StructB) String() string {
	s.val += 10
	fmt.Println("StructB String()")
	return "StructB:" + strconv.Itoa(s.val)
}

type Printable interface {
	String() string
}

func Println(p Printable) {
	fmt.Println(p.String())
}

func main() {
	a := &StructA{val: "AAA"}
	Println(a)

	b := &StructB{val: 10}
	Println(b)
}
