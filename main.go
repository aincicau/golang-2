package main

import (
	"fmt"
	"internship/golang2/entity"
)

const (
	Pi       = 3.14
	Timezone = "CET"
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

	/*var shape entity.Shape
	shape = entity.Circle{Radius: 2.4}
	circle, ok := shape.(entity.Rectangle)
	if !ok {
		fmt.Println("Its not a rectangle")
	}
	fmt.Println(circle.Area())*/

	/*names := make(map[string]int)
	names = map[string]int{
		"tom": 1,
	}

	value, _ := names["tom"]
	fmt.Println(value, len(names))

	for key, v := range names { //key in loc de index(ca la array)
		fmt.Println(key, v)
	}*/

	/*var n *int
	var m = new(int)
	fmt.Println(n, m)*/

	/*switch vehicle.(type) {
	case entity.Car:
		fmt.Println("This is a car")
	default:
		fmt.Println("default")
	}*/

}

func Test(a *int) {
	*a = 40
}

func GetCircumference(shape entity.Shape) {
	circle, _ := shape.(entity.Circle) //ok in loc de _
	fmt.Println(circle.Circumference())
}
