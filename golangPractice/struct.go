package main

import "fmt"

//type 宣告新型態
type Point struct {
	x float64
	y float64
}

func structTest() {
	p := Point{x: 3.0, y: 4.0}
	agent := struct {
		name string
		age  int
	}{name: "Jeff", age: 27}

	fmt.Println(p.x)
	fmt.Println(p.y)
	fmt.Println(agent.name)
	fmt.Println(agent.age)

}
