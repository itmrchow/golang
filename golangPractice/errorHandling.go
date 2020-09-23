package main

import (
	"errors"
	"fmt"
)

func errorHandling() {
	err := errors.New("some error")
	//err不為空
	if err != nil {
		fmt.Println(err)
	}

	//一般錯誤發生程式不停止，嚴重錯誤才會停
	fmt.Println("More message")
}
