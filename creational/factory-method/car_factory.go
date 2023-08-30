package factory_method

type CarFactory struct{}

func (c CarFactory) Produce() VehicleInterface {
	return Car{}
}
