package main

import "fmt"

func hello(name string) (bool, string) {
	fmt.Println("Hello " + name)

	return true, "success"
}
