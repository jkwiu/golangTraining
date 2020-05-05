package main

import "fmt"

func main() {
	// 마름모
	for i := 0; i < 4; i++ {
		for j := 4; j > i+1; j-- {
			fmt.Print(" ")
		}
		for k := 0; k < 2*i+1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := 0; i < 4; i++ {
		for l := 0; l < i+1; l++ {
			fmt.Print(" ")
		}
		for m := 6; m > 2*i+1; m-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
