package strtmethintr

import "math"

type Shape interface {
	Area() float64
}

type Rectange struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectange) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2.0)
}
