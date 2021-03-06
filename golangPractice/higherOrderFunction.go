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

//fold 將等長的串列組合成tuple
type Tuple1 struct {
	First  int
	Second int
}

func zip(mAry []int, nAry []int) func() (Tuple1, bool) {
	if len(mAry) != len(nAry) {
		panic("Unequal list length")
	}

	i := -1

	return func() (Tuple1, bool) {
		i++

		if i < len(mAry) {
			return Tuple1{First: mAry[i], Second: nAry[i]}, true
		}
		return Tuple1{First: 0, Second: 0}, false

	}
}

//列舉
type Tuple2 struct {
	Index   int
	Element int
}

func enumerate(arr []int) func() (Tuple2, bool) {
	i := -1

	return func() (Tuple2, bool) {
		i++
		if i < len(arr) {
			return Tuple2{Index: i, Element: arr[i]}, true
		}
		return Tuple2{Index: i, Element: 0}, false

	}

}
