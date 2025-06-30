package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	go greeter("Hello")       // This goroutine calls Done 5 times
	wg.Add(5)                 // ✅ Add BEFORE starting the goroutine
	greeter("World")          // Prints synchronously
	wg.Wait()                 // Waits for 5 Done calls
}

func greeter(s string) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}
