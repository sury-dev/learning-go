package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a rating between 1 through 5:")

	rating, _ := reader.ReadString('\n')

	numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Thanks for rating us with", numRating + 1, "stars!")
	}

}