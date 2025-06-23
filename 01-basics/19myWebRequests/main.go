package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.google.com/"

func main() {
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	//fmt.Println(response)

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(databytes))
}