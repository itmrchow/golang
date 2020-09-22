package main

import (
	"fmt"
	"log"
)

func arrayTest() {
	var langs [4]string
	langs[0] = "Go"
	langs[1] = "Python"
	langs[2] = "Ruby"
	langs[3] = "PHP"

	langs2 := [3]string{
		"Jeff",
		"Tony",
		"Amy",
	}

	if !(langs[0] == "Go") {
		log.Fatal("Wrong string")
	}

	for i, e := range langs2 {
		fmt.Println(fmt.Sprintf("%d: %s", i+1, e))
	}

	//不用索引
	for _, e := range langs2 {
		fmt.Println(e)
	}

}

func sliceTest() {
	lang := [3]string{
		"Jeff",
		"Tony",
		"Amy",
	}

	slice := lang[0:3]

	//slice儲存的是array的位子
	//修改slice會連同array的值一起修改
	slice[1] = "John"

	slice = append(slice, "James", "Lily")
	//移除用忽略的方式
	slice = append(slice[0:2], slice[4:len(slice)]...)

	for i, e := range slice {
		fmt.Println(fmt.Sprintf("%d:%s", i+1, e))
	}
}

func matrix() {
	matrix := [][]float64{
		[]float64{1, 2, 3},
		[]float64{4, 5, 6},
	}

	for _, e := range matrix {
		fmt.Println(e)
	}
}

func makeTest() {
	slice := make([]int, 5)

	for i := 0; i < len(slice); i++ {
		n := i + 1
		slice[i] = n * n
	}

	for _, v := range slice {
		fmt.Println(v)
	}
}
