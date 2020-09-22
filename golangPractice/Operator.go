package main

import (
	"fmt"
	"os"
	"runtime"
)

func operator() {
	assert((true && true) == true, "Wrong logic")
	assert((true && false) == false, "Wrong logic")
	assert((false && true) == false, "Wrong logic")
	assert((false && false) == false, "Wrong logic")

	assert((true || true) == true, "Wrong logic")
	assert((true || false) == true, "Wrong logic")
	assert((false || true) == true, "Wrong logic")
	assert((false || false) == false, "Wrong logic")

	assert((!true) == false, "Wrong logic")
	assert((!false) == true, "Wrong logic")
}

//檢查函式
func assert(cond bool, msg string) {
	//Caller 取得呼叫者資訊
	//f 檔案名稱
	//l 所在行數
	_, f, l, _ := runtime.Caller(1)

	if !cond {
		//Stderr標準錯誤
		fmt.Fprintf(os.Stderr, "Failed on (%s:%d):%s", f, l, msg)
		os.Exit(1)
	}

}
