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
	//out := make([]int, 0)

	for i, v := range arr {
		arr[i] = mapper(v)
	}
	return arr
}
