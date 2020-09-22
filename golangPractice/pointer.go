package main

import (
	"fmt"
	"log"
)

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
