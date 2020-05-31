package main

import (
	"strconv"
	"fmt"
)

type StructA struct {
	
}

func (a *StructA) AAA(x int) int {
	return x*x
}

func (a *StructA) BBB(x int) string{
	return "X= " + strconv.Itoa(x)
}

func main() {
	var c InterfaceA
	c = &StructA{}

	fmt.Println(c.BBB(3))
}