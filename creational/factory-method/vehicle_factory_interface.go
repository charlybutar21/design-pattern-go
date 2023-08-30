package factory_method

type VehicleFactoryInterface interface {
	Produce() VehicleInterface
}
