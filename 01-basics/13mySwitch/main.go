package main

import (
	"fmt"
	"math/rand"
)

func main() {

	num := rand.Intn(7)

	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
		fallthrough
	case 4:
		fmt.Println("Four")
		fallthrough // in this language we explixitely do the fallthrough operation
	case 5:
		fmt.Println("Five")
	case 6:
		fmt.Println("Six")
	default:
		fmt.Println("What the hell was that??")
	}
}
