package main

import "fmt"

func main() {
	getCircleArea()

	sayHello()
}

func getCircleArea() {
	const PI float64 = 3.1415927
	r := 10.0
	fmt.Println(r * r * PI)
}
