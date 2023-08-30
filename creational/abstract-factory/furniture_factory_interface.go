package abstract_factory

type FurnitureFactoryInterface interface {
	CreateChair() ChairInterface
	CreateTable() TableInterface
}
