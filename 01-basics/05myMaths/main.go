package main

import (
	"fmt"
	"math/big"
	//"math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Welcome to My Maths App")
	var num1 int = 10
	var num2 float32 = 20.5

	fmt.Println("Addition:", num1+int(num2)) // Convert float32 to int for addition //loss precision

	//random numbers

	//rand.Seed(42) // Seed the random number generator for reproducibility
	//randomNum := rand.Intn(100) // Generate a random number between 0 and 99
	
	//random number using crypto/rand
	randomNum2 , _ := rand.Int(rand.Reader, big.NewInt(100)) // Generate a random number between 0 and 99
	// The second return value is an error, which we ignore for simplicity
	
	fmt.Println("Random Number:", randomNum2)
}