package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	out := filter(arr, func(n int) bool {
		return n%2 == 0
	})

	fmt.Println(out)

}
