package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Web requests")
	PerformGetRequest("http://localhost:8000/get")
	PerformPostJsonRequest("http://localhost:8000/post")
	PerformPostFormRequest("http://localhost:8000/postform")
}

func PerformGetRequest(uri string){

	res, err := http.Get(uri)

	if err != nil{
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println("Status Code: ", res.StatusCode)
	fmt.Println("Content Length : ", res.ContentLength)
	// cont, _ := io.ReadAll(res.Body)
	
	// fmt.Println("Content : \n", string(cont))

	//Other Code

	var responseString strings.Builder
	cont, _ := io.ReadAll(res.Body)
	byteCount, _ := responseString.Write(cont)

	fmt.Println("Byte Count :", byteCount)
	fmt.Println(responseString.String())
}

func PerformPostJsonRequest(uri string){
	//fake json payload

	requestBody  := strings.NewReader(`
	{
		"name" : "jhgfkjg",
		"price" : 9,
		"platform" : "fgh.fgh"
	}`)

	res, err := http.Post(uri, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	cont, _ := io.ReadAll(res.Body)

	fmt.Println(string(cont))
}

func PerformPostFormRequest(myurl string)  {

	//formdata

	data := url.Values{}
	data.Add("firstname", "hitesh")
	data.Add("lastname", "choudhary")
	data.Add("email", "hitesh@go.dev")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}