package entity

type Bike struct {
	Wheels   int
	Drivable bool
}

func (bike Bike) NoOfWheels() int {
	return bike.Wheels
}

func (bike Bike) CanDrive() bool {
	return bike.Drivable
}
