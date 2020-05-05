package main

import "fmt"

func pop(c chan int)  {
  fmt.Println("Pop func")
  v := <-c
  fmt.Println(v)
}

func main()  {
  //채널 선언
  var c chan int
  //길이 1개의 채널 초기화
  c = make(chan int)

  go pop(c)
  c <- 10



}
