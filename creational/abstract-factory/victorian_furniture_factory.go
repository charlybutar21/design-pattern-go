package abstract_factory

type VictorianFurnitureFactory struct{}

func (f VictorianFurnitureFactory) CreateChair() ChairInterface {
	return VictorianChair{}
}

func (f VictorianFurnitureFactory) CreateTable() TableInterface {
	return VictorianTable{}
}
