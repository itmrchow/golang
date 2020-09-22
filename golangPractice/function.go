package main

import "fmt"

func hello(name string) (bool, string) {
	fmt.Println("Hello " + name)

	return true, "success"
}

func sum(args ...float64) float64 {
	sum := 0.0
	for _, v := range args {
		sum += v
	}

	return sum
}
