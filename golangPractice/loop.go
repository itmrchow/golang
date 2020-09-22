package main

import (
	"fmt"
)

//條件式loop
func loop() {
	i := 1

	for i <= 10 {
		fmt.Println(i)

		i++
	}
}

//計數器loop
func loop2() {
	for i := 1; i < 10; i++ {
		if 5 < i {
			break
		} else if 3 == i {
			continue
		} else if 2 == i {
			goto gotoInfo
		} else {
			fmt.Println(i)
		}
	}

	//標籤
gotoInfo:
	gotoInfo()
}

func gotoInfo() {
	fmt.Println("goto")
}
