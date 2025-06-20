package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your name:")

	var name string

	name, _ = reader.ReadString('\n') // ReadString reads until the first occurrence of the delimiter '\n'
	// The second return value is an error, which we ignore for simplicity
	fmt.Println("Hello,", name)


}