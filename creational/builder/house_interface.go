package builder

type HouseInterface interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}
