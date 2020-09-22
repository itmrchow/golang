package main

import (
	"fmt"
)

//條件式loop
func loop() {
	i := 1

	for i <= 10 {
		fmt.Println(i)

		i++
	}
}

//計數器loop
func loop2() {
	for i := 1; i < 10; i++ {
		if 5 < i {
			break
		}
		fmt.Println(i)

	}
}
