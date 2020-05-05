package main

import "fmt"

func main() {
	a := 4
	b := 2

	fmt.Printf("%v\n", a&b)
	fmt.Printf("%v\n", a|b)
	fmt.Printf("%v\n", a^b) // XOR연산 ^b 이면 NOT연산

	c := 21
	d := c % 10
	c = c / 10
	e := c % 10

	fmt.Printf("첫번째 수: %v 두번째 수: %v\n", d, e) // 자리수의 숫자를 뽑아낼 수 있다

	f := 4

	fmt.Println(f << 1)
	fmt.Println(f >> 1)

	var g bool
	g = 3 > 4

	fmt.Println(g)
}
