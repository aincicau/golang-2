package entity

type Vehicle interface {
	NoOfWheels() int
	CanDrive() bool
}

type Tyre interface {
	Material()
}

type Train interface {
	Vehicle
	Tyre
}

type ExpressTrain struct {
	Wheels int
}