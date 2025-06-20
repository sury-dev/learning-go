package main

import (
	"fmt"
	"sort"
)

func main() {
	var slice1 []string
	slice1 = append(slice1, "Hello", "World")
	slice1 = append(slice1, "from")
	slice1 = append(slice1, "Go")
	slice1 = append(slice1, "Lang")

	fmt.Println("Slice 1:", slice1)
	fmt.Println("Length of slice 1:", len(slice1))
	fmt.Printf("Type of slice 1: %T\n", slice1)

	fmt.Println("Slice 1 from index 1 to 3:", slice1[1:4])
	fmt.Println("Slice 1 from index 2 to end:", slice1[2:])
	fmt.Println("Slice 1 from start to index 3:", slice1[:4])

	var slice2 = []string{"I", "am", "a", "Go", "developer"}
	fmt.Println("Slice 2:", slice2)

	scores := []int{10, 20, 30, 40, 50}
	fmt.Println("Scores:", scores)
	scores[2] = 100 // Change the value at index 2
	fmt.Println("Updated Scores:", scores)

	scores = append(scores, 60, 70) // Append new values to the slice
	fmt.Println("Scores after appending:", scores)

	//Operations on slices

	sort.Ints(scores) // Sort the scores slice
	fmt.Println("Sorted Scores:", scores)

	fmt.Println("Is the slice sorted", sort.IntsAreSorted(scores))

	//Removing a value based on index

	scores = append(scores[:6], scores[7:]...)
	fmt.Println("Slice after slicing :", scores)
}
