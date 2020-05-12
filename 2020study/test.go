package main

import "fmt"

func main() {
	e := make([]int, 2, 5)

	fmt.Println(cap(e), len(e))

}
