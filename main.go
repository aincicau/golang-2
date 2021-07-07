package main

import (
	"fmt"
	"internship/golang2/entity"
)

func main() {
	/*var a = 10
	var b *int
	b = &a
	*b = 20
	Test(b)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)*/

	/*var vehicle entity.Vehicle
	vehicle = entity.Car{Wheels: 4}
	fmt.Println(vehicle)*/

	var shape entity.Shape
	shape = entity.Circle{Radius: 2.4}
	circle, ok := shape.(entity.Rectangle)
	if !ok {
		fmt.Println("Its not a rectangle")
	}
	fmt.Println(circle.Area())
}

func Test(a *int) {
	*a = 40
}

func GetCircumference(shape entity.Shape) {
	circle, ok := shape.(entity.Circle)
	fmt.Println(circle.Circumference())
}
