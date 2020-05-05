package main

import "fmt"

func main() {
	//*
	//**
	//***
	//**
	//*
	for i := 0; i < 3; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

	for i := 0; i < 3; i++ {
		for j := 2; j > i; j-- {
			fmt.Print("*")
		}
		fmt.Println("")
	}

}
