package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Welcome to My Time App")

	//get input
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your name:")
	var name string
	
	presentTime := time.Now() // Get the current time
	fmt.Println("Current Time:", presentTime)
	fmt.Printf("Type of presentTime is: %T", presentTime)
	
	//formatted time
	formattedTime := presentTime.Format("01-02-2006 15:04:05 Monday") // Format the time in a specific layout
	// The layout must be in the reference time Mon Jan 2 15:04:05 MST 2006
	// restrictions 
	// 01 - month, 02 - day, 2006 - year, 15 - hour in 24-hour format, 04 - minute, 05 - second, Monday - day of the week
	fmt.Println("\nFormatted Time:", formattedTime)
	
	//manual date creation
	manualTime := time.Date(2023, time.October, 10, 12, 0, 0, 0, time.UTC) // Create a specific date and time
	fmt.Println("Manual Time:", manualTime.Format("01/02/2006 15:04: Monday"))

	
	name, _ = reader.ReadString('\n') // ReadString reads until the first occurrence of the delimiter '\n'
	// The second return value is an error, which we ignore for simplicity
	fmt.Println("Hello,", name)
}