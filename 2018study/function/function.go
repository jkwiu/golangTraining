package main

import "fmt"

/*
func add(x int, y int) int { //input 두개, x,y output 1개
	return x + y
}

func ()  {

}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d + %d = %d\n", i, i+2, add(i, i+2))
	}
}


func main() {
	a, b := fun1(2, 3)
	fmt.Println(a, b)
}

func fun1(x, y int) (int, int) {
	fun2(x, y)
	return y, x
}

func fun2(x, y int) {
	fmt.Println("func2", x, y)
}
*/

//재귀호출
func main() {
	s := sum(10, 0)
	fmt.Println(s)
}

func sum(x int, s int) int {
	if x == 0 {
		return s
	}
	s += x
	return sum(x-1, s)
}
