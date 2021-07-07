package entity

type Rectangle struct {
	Width  float32
	Length float32
}

func (rectangle Rectangle) Area() float32 {
	return rectangle.Length * rectangle.Width
}
