package factory_method

type Car struct{}

func (c Car) GetName() string {
	return "Mobil"
}
