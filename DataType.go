package main

import "fmt"

func main() {
	getCircleArea()
}

func sayHello() {
	var msg string
	msg = "Hello World"

	name := "Jeff"

	fmt.Println(msg)
	fmt.Println(name)
}

func getCircleArea() {
	const PI float64 = 3.1415927
	r := 10.0
	fmt.Println(r * r * PI)
}
