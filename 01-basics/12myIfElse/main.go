package main

import "fmt"

func main() {

	value := 76

	if value == 0 {
		fmt.Println("Zero")
	} else if value > 0 {
		fmt.Println("Positive")
	} else {
		fmt.Println("Negative")
	}

	if num := -3 ; num < 0 {
		fmt.Println("Negative")
	}

}