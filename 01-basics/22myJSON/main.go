package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"courseName"`
	Price    int
	Platform string
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome \n ------------------")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"GoLang Mastery", 349, "udemy.com", "def456", []string{"backend", "go", "api"}},
		{"UI/UX Design Principles", 199, "coursera.org", "ghi789", []string{"design", "figma", "uiux"}},
		{"Fullstack Web Development", 399, "edx.org", "jkl012", []string{"frontend", "backend", "javascript", "node"}},
		{"Python for Data Science", 249, "datacamp.com", "mno345", []string{"python", "data-science", "machine-learning"}},
		{"DevOps with Docker & Kubernetes", 299, "pluralsight.com", "pqr678", []string{"devops", "docker", "k8s"}},
		{"TypeScript Essentials", 189, "frontendmasters.com", "stu901", nil},
	}

	//package data as json data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {
	jsonData := []byte(`
	{
                "courseName": "GoLang Mastery",
                "Price": 349,
                "Platform": "udemy.com",
                "tags": [
                        "backend",
                        "go",
                        "api"
                ]
        }`)

	var lcoCourse course

	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonData, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("The JSON data was not valid")
	}

	//some cases where you just want to add data to key value

	var myOnlineData map[string]any
	json.Unmarshal(jsonData, &myOnlineData)
	fmt.Println(myOnlineData)

	for k, v := range myOnlineData {
		fmt.Println(k, ":", v)
	}

}
