package main

import (
	"fmt"
	"golangTraining/2020study/ch37Map"
)

func main() {
	fmt.Println("abcde = ", ch37Map.Hash("abcde"))
	fmt.Println("abcde = ", ch37Map.Hash("abcde"))
	fmt.Println("tbcde = ", ch37Map.Hash("tbcde"))
	fmt.Println("abcdf = ", ch37Map.Hash("abcdf"))
	fmt.Println("abcasdfdfasdf = ", ch37Map.Hash("abcasdfdfasdf"))

	m := ch37Map.CreateMap()
	m.Add("AAA", "0107777777")
	m.Add("BBB", "0108888888")
	m.Add("CDFEFDFDFDFFD", "0101111111")
	m.Add("CCC", "01712387842")

	fmt.Println("AAA = ", m.Get("AAA"))
	fmt.Println("BBB = ", m.Get("BBB"))
	fmt.Println("CDFEFDFDFDFFD = ", m.Get("CDFEFDFDFDFFD"))
	fmt.Println("CCC", m.Get("CCC"))
}
