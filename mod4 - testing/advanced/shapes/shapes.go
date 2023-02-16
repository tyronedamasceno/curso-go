package shapes

import (
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Heigth float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Heigth * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}
