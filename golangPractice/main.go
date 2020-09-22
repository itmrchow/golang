package main

import "log"

func main() {
	p := NewPoint(3, 4)

	if !(p.X == 3.0) {
		log.Fatal("Wrong value")
	}

	if !(p.Y == 4.0) {
		log.Fatal("Wrong value")
	}
}
