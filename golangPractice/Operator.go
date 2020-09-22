package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
)

func operator() {
	assert(4+3 == 7, "4 + 3 should be 7")
	assert(4-3 == 1, "4 - 3 should be 1")
	assert(4*3 == 12, "4 * 3 should be 12")

	assert(4/3 == 1, "4 / 3 should be 1")
	//Abs 返回絕對值
	assert(math.Abs(4.0/3.0-1.333333) < 0.00001, "4.0 / 3.0 should be 1.333333")

	assert(4%3 == 1, "4 % 3 should be 1")
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
