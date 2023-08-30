package factory_method

type Motorcycle struct{}

func (m Motorcycle) GetName() string {
	return "Sepeda Motor"
}
