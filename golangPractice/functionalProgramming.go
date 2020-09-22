package main

import "fmt"

func funcTest(msg string) {
	fmt.Println(msg)
}

//方法作為參數
func apply(arr []float64, callback func(float64) float64) []float64 {
	out := make([]float64, len(arr))

	for i, v := range arr {
		out[i] = callback(v)
	}

	return out
}

//閉包
func num() func() int {
	n := -1

	return func() int {
		n += 1
		return n
	}
}
