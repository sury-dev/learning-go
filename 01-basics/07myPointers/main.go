package main

import "fmt"

func main() {
	fmt.Println("Pointers in Go")

	myNumber := 23

	fmt.Println("My number:", myNumber)
	fmt.Println("Address of my number:", &myNumber)

	//declaring a pointer
	//var myPointer *int;
	//myPointer = &myNumber // myPointer now holds the address of myNumber
	//or
	myPointer := &myNumber // myPointer now holds the address of myNumber


	fmt.Println("Value of myPointer:", myPointer)
	fmt.Println("Value at the address of myPointer:", *myPointer)

	*myPointer = 42 // changing the value at the address of myPointer
	fmt.Println("New value of myNumber:", myNumber)
}