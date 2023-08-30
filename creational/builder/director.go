package builder

type Director struct {
	builder HouseInterface
}

func NewDirector(b HouseInterface) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) SetBuilder(b HouseInterface) {
	d.builder = b
}

func (d *Director) BuildHouse() House {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()

	return d.builder.getHouse()
}
