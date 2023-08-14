package main

import "fmt"

func main() {
	test()
	consolePrint()
	consolePrint1()
}

func test() {
	fmt.Println("Hello, World!")
}
func consolePrint() {
	fmt.Println("Hi! This is Print Function")
}

func consolePrint1() {
	fmt.Println("Hi! This is Print1 Function")
	test()
}
