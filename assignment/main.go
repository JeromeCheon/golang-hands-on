package main

import (
	"fmt"
	"math"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return math.Pow(s.sideLength, 2)
}

func main() {
	t := triangle{
		height: float64(60),
		base:   float64(40),
	}

	s := square{
		sideLength: float64(30),
	}
	printArea(t)
	printArea(s)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
