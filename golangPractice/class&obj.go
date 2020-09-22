package main

type PointByClass struct {
	X float64
	Y float64
}

type Vector []float64

//使用方法作為建構函數
func NewPoint(x float64, y float64) *PointByClass {
	p := new(PointByClass)
	p.X = x
	p.Y = y

	return p
}

func NewVector(args ...float64) Vector {
	return args
}

func WithSize(s int) Vector {
	v := make([]float64, s)

	return v
}
