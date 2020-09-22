package main

import (
	"fmt"
	"log"
)

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
