package main

import "fmt"

// 관계만 선언
type Jam interface {
	GetOneSpoon() SpoonOfJam
}

type Bread struct {
	val string
}

type StrawberryJam struct {
}

type OrangeJam struct {
}

type AppleJam struct {
}

type SpoonOfJam interface {
	String() string
}

type SpoonOfStrawberryJam struct {
}

type SpoonOfOrangeJam struct {
}

type SpoonOfAppleJam struct {
}

func (s *SpoonOfAppleJam) String() string {
	return "+ Apple"
}

func (s *SpoonOfOrangeJam) String() string {
	return "+ Orange"
}

func (s *SpoonOfStrawberryJam) String() string {
	return "+ strawberry"
}

func (j *StrawberryJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfStrawberryJam{}
}

func (j *OrangeJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfOrangeJam{}
}

func (j *AppleJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfAppleJam{}
}

func (b *Bread) PutJam(jam Jam) {
	spoon := jam.GetOneSpoon()
	b.val += spoon.String()
}

func (b *Bread) String() string {
	return "bread " + b.val
}

func main() {
	bread1 := &Bread{}
	bread2 := &Bread{}
	bread3 := &Bread{}
	jam1 := &StrawberryJam{}
	jam2 := &OrangeJam{}
	jam3 := &AppleJam{}
	bread1.PutJam(jam1)
	bread2.PutJam(jam2)
	bread3.PutJam(jam3)

	fmt.Println(bread1)
	fmt.Println(bread2)
	fmt.Println(bread3)
}
