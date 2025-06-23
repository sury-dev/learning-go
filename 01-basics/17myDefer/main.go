package main

import "fmt"

func main() {
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	fmt.Println("Hello Defer")
	mydefer()
}

func mydefer(){
	defer fmt.Println()
	for i := 0; i < 6; i++ {
		defer fmt.Print(i, " ")
	}
}