package main

import "math"

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

func Dist(p1 *PointByClass, p2 *PointByClass) float64 {
	xSqr := math.Pow(p1.GetX()-p2.GetX(), 2)
	ySqr := math.Pow(p1.GetY()-p2.GetY(), 2)

	return math.Sqrt(xSqr + ySqr) /* 32 */
}

func (p *PointByClass) area() float64 {
	return (p.x * p.y) / 2
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

type Point3D struct {
	*PointByClass
	z float64
}

func NewPoint3D(x float64, y float64, z float64) *Point3D {
	p := new(Point3D)
	p.PointByClass = NewPoint(x, y)

	// p.SetX(x)
	// p.SetY(y)
	p.SetZ(z)

	return p
}

func (p *Point3D) SetZ(z float64) {
	p.z = z
}

func (p *Point3D) GetZ() float64 {
	return p.z
}
