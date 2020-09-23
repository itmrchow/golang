package main

//filter 傳入一串列與函式條件，回傳相符的串列
func filter(arr []int, predicate func(int) bool) []int {
	out := make([]int, 0)

	for _, v := range arr {
		if predicate(v) {
			out = append(out, v)
		}
	}
	return out
}

//map 讓每個元素經過函式轉換
func mapply(arr []int, mapper func(int) int) []int {
	out := make([]int, len(arr))

	for i, v := range out {
		out[i] = mapper(v)
	}
	return out
}

//reduce 處理後傳入下一輪
func reduce(arr []int, reducer func(int, int) int) int {
	if len(arr) == 0 {
		return 0
	} else if len(arr) == 1 {
		return arr[0]
	}

	n := arr[0]
	for i := 1; i < len(arr); i++ {
		n = reducer(n, arr[i])
	}

	return n
}

//parition 依規則分類
func partition(arr []int, predicate func(int) bool) ([]int, []int) {
	out1 := make([]int, 0)
	out2 := make([]int, 0)

	for _, v := range arr {
		if predicate(v) {
			out1 = append(out1, v)
		} else {
			out2 = append(out2, v)
		}
	}

	return out1, out2
}
