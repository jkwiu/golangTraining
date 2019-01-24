package main

import "fmt"

func main() {
	/*
		for dan := 1; dan <= 9; dan++ {
			fmt.Printf("%d단\n", dan)

			for j := 1; j <= 9; j++ {
				fmt.Printf("%d * %d = %d\n", dan, j, dan*j)
			}

			fmt.Println()
		}
	*/

	for i := 0; i < 3; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2-i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
