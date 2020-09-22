package main

import "fmt"

//Gender 性別
type Gender int

//iota 常數計數器
const (
	Male Gender = iota
	Female
)

type Person struct {
	name   string
	gender Gender
	age    uint
}

func structTest() {
	me := Person{gender: Male, age: 27, name: "Jeff"}

	fmt.Println(me.name)
	fmt.Println(me.age)

	if me.gender == Male {
		fmt.Println("Male")
	} else {
		fmt.Println("Female")
	}

}
