package main

import (
	"fmt"
)

func errorHandling() {
	panic("Some error")

	// It didn't occur.
	fmt.Println("More message")
}
