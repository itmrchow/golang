package main

import (
	"fmt"
	"log"
)

type Point struct {
	x float64
	y float64
}

//靜態配置
func pointerTest() {
	n := 2

	//變數位址
	nPtr := &n

	fmt.Println(nPtr)

	//取位址值
	if !(*nPtr == 2) {
		log.Fatal("Wrong value")
	}

}

//動態配置
func pointerTest2() {
	//變數位址
	nPtr := new(int)
	*nPtr = 2

	fmt.Println(nPtr)

	//取位址值
	if !(*nPtr == 2) {
		log.Fatal("Wrong value")
	}

}

//動態配置struct記憶體位址
func pointerTest3() {
	p := new(Point)

	p.x = 3.0
	p.y = 4.0

	fmt.Println(fmt.Sprintf("(%.2f, %.2f)", p.x, p.y))
}
