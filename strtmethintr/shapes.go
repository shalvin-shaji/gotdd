package strtmethintr

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Breadth float64
	Height  float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Breadth * t.Height
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2.0)
}
