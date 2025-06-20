package main

import "fmt"

const LoginToken string = "abc123" // LoginToken is a constant string that can be used throughout the program
//First letter of a constant is capitalized, so it is exported and can be used in other packages

func main() {
	// Declare a variable
	var username string = "JohnDoe"
	fmt.Println("Value :", username)
	fmt.Printf("Username is of type %T \n", username)

	var isLoggedIn bool = true
	fmt.Println("Value :", isLoggedIn)
	fmt.Printf("Variable is of type %T \n", isLoggedIn)

	var smallInt uint8 = 255 // uint8 can hold values from 0 to 255
	fmt.Println("Value :", smallInt)
	fmt.Printf("Variable is of type %T \n", smallInt)

	//small float
	var smallFloat float32 = 3.14998809808098098809890 // float32 can hold up to 7 decimal places
	fmt.Println("Value :", smallFloat)
	fmt.Printf("Variable is of type %T \n", smallFloat)

	//default values and some aliases
	var defaultInt int // default value is 0
	fmt.Println("Value :", defaultInt)
	fmt.Printf("Variable is of type %T \n", defaultInt)

	//implicit way
	var implicitInt = 42 // type is inferred as int
	fmt.Println("Value :", implicitInt)
	fmt.Printf("Variable is of type %T \n", implicitInt)

	//no var style

	implicitString := "Hello, World!" // type is inferred as string
	// only works inside functions
	// cannot use := outside of functions
	fmt.Println("Value :", implicitString)
	fmt.Printf("Variable is of type %T \n", implicitString)

	// using constants
	fmt.Println("Login Token:", LoginToken)
	fmt.Printf("LoginToken is of type %T \n", LoginToken)
}