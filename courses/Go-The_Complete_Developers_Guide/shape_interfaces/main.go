package main

import (
	"fmt"
	"math"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}
type square struct {
	sideLength float64
}

func main() {
	s := square{sideLength: 5.7}
	t := triangle{
		base:   5.7,
		height: 5.7,
	}

	printArea(s)
	printArea(t)
}

func (s square) getArea() float64 {
	return math.Pow(s.sideLength, 2)
}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}
