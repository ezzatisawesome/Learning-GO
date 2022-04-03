package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rect struct {
	width  float64
	height float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func main() {
	c1 := circle{4.5}
	r1 := rect{5, 8.2}
	shapes := []shape{c1, r1}
	for _, shape := range shapes {
		fmt.Println(shape.area())
	}
}
