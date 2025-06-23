package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello Files")
	content := "This need to go in the file - Copy Paste"

	file, err := os.Create("myFile.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}
	readMyFile("./myFile.txt")

	fmt.Println("Length : ", length)
}

func readMyFile(filename string)  {
	databyte, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	fmt.Println("Data in the file is : \n", string(databyte))
}