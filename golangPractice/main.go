package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	filterOut := filter(arr, func(n int) bool {
		return n%2 == 0
	})

	mapOut := mapply(arr, func(n int) int {
		return n * 5
	})

	reduceOut := reduce(arr, func(n1 int, n2 int) int {
		return n1 + n2
	})

	fmt.Printf("filterOut:%v\n", filterOut)
	fmt.Printf("mapOut:%v\n", mapOut)
	fmt.Printf("reduceOut:%v\n", reduceOut)

}
