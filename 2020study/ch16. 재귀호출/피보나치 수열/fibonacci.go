package main

import "fmt"

// 피보나치 수열
// f(0) = 1
// f(1) = 1
// 2 이상일 때 부터
// f(x) = f(x-1) + f(x-2)
func main() {
	rst := f(10)
	fmt.Println(rst)
}

func f(x int) int {
	if x == 0 {
		return 1
	}
	if x == 1 {
		return 1
	}
	return f(x-1) + f(x-2)
}
