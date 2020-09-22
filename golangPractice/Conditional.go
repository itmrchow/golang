package main

import (
	"fmt"
	"math/rand"
	"time"
)

func confitional() {
	//用系統時間來取得亂數
	//UnixNano 返回unix時間
	s := rand.NewSource(time.Now().UnixNano())

	r := rand.New(s)

	n := r.Intn(10) + 1

	if n%2 == 0 {
		fmt.Printf("%d is even\n", n)
	} else {
		fmt.Printf("%d is odd\n", n)
	}
}

func confitional2() {
	now := time.Now()

	weekOfDay := now.Weekday()

	switch weekOfDay {
	case 6, 7:
		fmt.Println("Weekend")
	case 5:
		fmt.Println("Thank God. It's Friday")
	case 3:
		fmt.Println("Hump day")
	default:
		fmt.Println("Week")
	}

}
