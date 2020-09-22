package main

import (
	"fmt"
	"strconv"
)

type IAnimal interface {
	Speak()
}

type AnimalType int

const (
	Duck AnimalType = iota
	Dog
	Tiger
)

type DuckClass struct {
}
type DogClass struct {
}
type TigerClass struct {
}

func NewDuck() *DuckClass {
	return new(DuckClass)
}

func NewDog() *DogClass {
	return new(DogClass)
}

func NewTiger() *TigerClass {
	return new(TigerClass)
}

func (d *DogClass) Speak() {
	fmt.Println("旺")
}

func (d *DuckClass) Speak() {
	fmt.Println("呱")
}

func (d *TigerClass) Speak() {
	fmt.Println("Halum")
}

func New(t AnimalType) IAnimal {
	switch t {
	case Duck:
		return NewDuck()
	case Dog:
		return NewDog()
	case Tiger:
		return NewTiger()
	default:
		//恐慌中斷
		panic("Unknown animal type")
	}
}

type Vector []float64

func NewVector(args ...float64) Vector {
	return args
}

func (v Vector) String() string {
	out := "("

	for i, e := range v {
		out += strconv.FormatFloat(e, 'f', -2, 64)

		if i < len(v)-1 {
			out += ", "
		}
	}

	out += ")"

	return out
}
