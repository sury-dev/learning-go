package main

import "fmt"

func main() {
	nums := [5]int{1, 5, 8, 6, 8}

	for index := range nums {
		fmt.Println(index)
	}


	for _ , value := range nums{
		fmt.Println(value)
	}

	for i := 0; i<2 ; i++ {
		fmt.Println("Nums")
	}

	i := 6
	for i < 10 {
		fmt.Println(i)
		i++
	}
}