package builder

type IglooHouse struct {
	House
}

func NewIglooHouse() *IglooHouse {
	return &IglooHouse{}
}

func (i *IglooHouse) setWindowType() {
	i.WindowType = "Snow Window"
}

func (i *IglooHouse) setDoorType() {
	i.DoorType = "Snow Door"
}

func (i *IglooHouse) setNumFloor() {
	i.Floor = 1
}

func (i *IglooHouse) getHouse() House {
	return i.House
}
