package entity

import "math"

type Circle struct {
	Radius float32
}

func (circle Circle) Area() float32 {
	return math.Pi * (circle.Radius) * (circle.Radius)
}

func (circle Circle) Circumference() float32 {
	return 2 * math.Pi * circle.Radius
}
