package main

import (
	"fmt"
	"log"
	"net/http"

	routers "github.com/suryansh/mongoapi/Routers"
)

func main() {
	fmt.Println("Mongo API")
	fmt.Println("Server is starting")
	fmt.Println("Server is listening on : https://localhost:4000")
	log.Fatal(http.ListenAndServe(":4000", routers.Router()))
}
