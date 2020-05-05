package main

import "fmt"

func main() {
	// 피라미드

	for i := 0; i < 4; i++ {
		for j := 4; j > i; j-- {
			fmt.Print(" ")
		}
		for k := 0; k < 2*i+1; k++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
