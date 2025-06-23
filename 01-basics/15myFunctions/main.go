package main

import "fmt"

func main() {
	greeter()
	fmt.Println(adder(23,65))
	ans, m := proAdder(3,4,5,6,1,5,2,4,3,2,2)
	fmt.Println("ans : ", ans, "Message:", m)
}
func greeter() {
	fmt.Println("Hey You! Stop right there.")
}
func adder(num1 int, num2 int) int {
	return num1 + num2;
}
func proAdder(num ...int) (int, string) {
	ans := 0
	for _ , v := range num{
		ans += v
	}
	return ans, "All Added"
}