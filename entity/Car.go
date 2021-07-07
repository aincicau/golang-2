package entity

type Car struct {
	Wheels int
}

func (car Car) NoOfWheels() int {
	return car.Wheels
}

func (car Car) CanDrive() bool {
	return true
}
