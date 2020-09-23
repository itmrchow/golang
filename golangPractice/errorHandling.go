package main

import (
	"fmt"
)

func errorHandling() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Recover from panic")
		}
	}()

	panic("some error")

	fmt.Println("More message")
}
