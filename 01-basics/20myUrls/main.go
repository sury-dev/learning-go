package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://www.sury-dev.com:6556/profile/projects?project_id=34567&traffic=1234"

func main() {
	result, _ := url.Parse(myUrl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Printf("Type of query params are %T\n", qparams)

	for key, value := range qparams {
		fmt.Println("Key :", key, "; Value :", value)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawQuery: "user=me",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println(anotherUrl)
}
