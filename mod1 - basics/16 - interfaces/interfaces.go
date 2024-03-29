package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

func writeArea(s shape) {
	fmt.Printf("shape area is %0.2f\n", s.area())
}

type rectangle struct {
	heigth float64
	width  float64
}

func (r rectangle) area() float64 {
	return r.heigth * r.width
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func main() {
	r := rectangle{10, 15}
	c := circle{10}
	writeArea(r)
	writeArea(c)
}
