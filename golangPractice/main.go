package main

import (
	"fmt"
	"log"
)

func main() {
	p := NewPoint(0, 0)
	p2 := NewPoint(3, 5)

	if !(p.x == 0) {
		log.Fatal("Wrong value")
	}

	if !(p.y == 0) {
		log.Fatal("Wrong value")
	}

	p.SetX(3)
	p.SetY(4)

	fmt.Println(p.GetX(), p.GetY())

	fmt.Println(p.area())
	fmt.Println(Dist(p, p2))

}
