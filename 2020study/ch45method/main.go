package main

import "fmt"

func main() {
	bread := &Bread{val: "bread"}
	jam := &Jam{}

	bread.PutJam(jam)
	fmt.Println(bread) 
}

type Bread struct {
	val string
}
type Jam struct {
	
}

func (b *Bread) PutJam(jam *Jam){
	b.val += jam.GetVal()
}

func (b *Bread) String() string{	// String()가 정의되어 있으면 결과 값을 출력한다.
	return b.val
}

func (j *Jam) GetVal() string{
	return " + jam"
}