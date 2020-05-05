package main

import "fmt"

type Student struct {
	name string
	age  int

	grade string
	class string
}

func (s *Student) PrintSungjuk() { //*Student의 변수의 메모리 주소를 가르키는 s
	fmt.Println(s.class, s.grade)
}

func (s *Student) InputSungjuk(class string, grade string) {
	s.class = class
	s.grade = grade
}

func main() {
	var s Student = Student{name: "jk", age: 33, class: "수학", grade: "A+"}

	s.InputSungjuk("과학", "A++")
	s.PrintSungjuk()

}
