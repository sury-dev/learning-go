package main

import "fmt"

func main() {
	fmt.Println("Pointers in Go")

	var arr [5]string
	arr[0] = "Hello"
	arr[1] = "World"
	arr[2] = "from"
	// arr[3] = "Go"
	// arr[4] = "Lang"

	fmt.Println("Array:", arr)
	fmt.Println("Length of array:", len(arr))
	fmt.Println("Capacity of array:", cap(arr))

	var arr2 = [5]string{"I", "am", "a", "Go", "developer"}

	arr3 := [5]string{"Learning", "Go", "is", "fun", "!"}
	fmt.Println("Array 3:", arr3)

	fmt.Println("Array 2:", arr2)
}