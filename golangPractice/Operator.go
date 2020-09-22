package main

import (
	"fmt"
	"os"
	"runtime"
)

func operator() {
	assert(4 == 4, "4 should be equal to 4")
	assert(4 != 3, "4 should not be equal to 3")

	assert(4 > 3, "4 should be greater than 3")
	assert(4 >= 3, "4 should be greater than or equal to 3")

	assert(4 < 5, "4 should be less than 5")
	assert(4 <= 5, "4 should be less than or equal to 5")
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
