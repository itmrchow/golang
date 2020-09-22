package main

type PointByClass struct {
	x float64
	y float64
}

//使用方法作為建構函數
func NewPoint(x float64, y float64) *PointByClass {
	p := new(PointByClass)
	p.x = x
	p.y = y

	return p
}

//寫入需要用pointer
func (p *PointByClass) GetX() float64 {
	return p.x
}

func (p *PointByClass) GetY() float64 {
	return p.y
}

func (p *PointByClass) SetX(x float64) {
	p.x = x
}

func (p *PointByClass) SetY(y float64) {
	p.y = y
}
