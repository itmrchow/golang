package main

import (
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

}
