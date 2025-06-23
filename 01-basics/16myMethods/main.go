package main

import "fmt"

func main() {
	fmt.Println("welcome to my struct class")
	//no inheritance in go lang

	surydev := User{"Sury-Dev", "sury.dev@coderhood.com", true, 19}

	fmt.Println("Details of the user :", surydev)
	fmt.Printf("Detailed : %+v \n", surydev)

	surydev.getStatus()
}

type User struct {
	name   string
	mail   string
	status bool
	id     int
}

func (u User) getStatus() {
	fmt.Println(u.status)
}
