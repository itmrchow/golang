package main

import "fmt"

func main() {
	f := funcTest

	f("hihi")

	arr := []float64{1, 2, 3}

	//平方
	sqr := apply(arr, func(n float64) float64 {
		return n * n
	})

	//立方
	cub := apply(arr, func(n float64) float64 {
		return n * n * n
	})

	fmt.Printf("平方：%f", sqr)
	fmt.Printf("立方：%f", cub)
}
