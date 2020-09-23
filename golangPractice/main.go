package main

import "fmt"

func main() {
	pings := make(chan string, 2)
	pongs := make(chan string, 2)

	ping(pings, "passed message")
	ping(pings, "passed message2")
	pong(pings, pongs)
	pong(pings, pongs)
	fmt.Println(<-pongs)
	fmt.Println(<-pongs)

	close(pongs)
	close(pongs)

}
