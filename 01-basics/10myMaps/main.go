package main

import "fmt"

func main() {

	fmt.Println("File on maps")

	lang := make(map[string]string)
	
	lang["js"] = "javascript"
	lang["cpp"] = "C++"
	lang["go"] = "goLang"

	fmt.Println("List of all languages", lang)
	fmt.Println("Full Form of JS", lang["js"])

	delete(lang, "js") // to detelete this entry

	fmt.Println("Lang after deletion", lang)

	//for loop

	for key, value := range lang{
		fmt.Println("Key is", key, "and value is", value)
	}
}