//OOP프로그래밍

package main

import "fmt"

type Bread struct {
  val string
}

type Jam struct {

}

func (b *Bread) PutJam(jam *Jam) {
  b.val += jam.GetVal()
}

func (j *Jam) GetVal() string {
  return "+Jam"
}

func (b *Bread) String() string {
  return b.val
}

func main()  {
  bread := &Bread{val:"bread"}
  jam := &Jam{}

  bread.PutJam(jam)

  fmt.Println(bread)
}
