package main

import (
	"fmt"
	"os"
	"runtime"
)

func operator() {
	/* 3 is 0011
	   5 is 0101 */

	/*  0011
	 &) 0101
	---------
	    0001  */
	assert((3&5) == 1, "3 & 5 should be 1")

	/*  0011
	 |) 0101
	---------
		0111  */
	assert((3|5) == 7, "3 | 5 should be 7")

	/*  0011
	 ^) 0101
	---------
		0110  */
	assert((3^5) == 6, "3 ^ 5 should be 6")

	/* <<) 0000 0101
	---------------
		   0000 1010  */
	assert((5<<1) == 10, "5 << 1 should be 10")

	/* >>) 0000 0101
	---------------
		   0000 0010  */
	assert((5>>1) == 2, "5 >> 1 should be 2")
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
