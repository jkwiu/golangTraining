package main

import "fmt"

//a[1:]읋 하면 가르키는 위치가 계속 달라져서 되는 거임
func RemoveBack(a []int) ([]int, int) {
	return a[:len(a)-1], a[len(a)-1]
}

func RemoveFront(a []int) ([]int, int) {
	return a[1:], a[0]
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < 5; i++ {
		var back int
		a, back = RemoveFront(a)
		fmt.Printf("%d", back)
	}

	fmt.Println(a)
}
