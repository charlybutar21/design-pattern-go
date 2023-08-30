package abstract_factory

type ModernFurnitureFactory struct{}

func (f ModernFurnitureFactory) CreateChair() ChairInterface {
	return ModernChair{}
}

func (f ModernFurnitureFactory) CreateTable() TableInterface {
	return ModernTable{}
}
