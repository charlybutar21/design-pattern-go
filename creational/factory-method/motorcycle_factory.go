package factory_method

type MotorcycleFactory struct{}

func (m MotorcycleFactory) Produce() VehicleInterface {
	return Motorcycle{}
}
