package builder

type NormalHouse struct {
	House
}

func NewNormalHouse() *NormalHouse {
	return &NormalHouse{}
}

func (n *NormalHouse) setWindowType() {
	n.WindowType = "Wooden Window"
}

func (n *NormalHouse) setDoorType() {
	n.DoorType = "Wooden Door"
}

func (n *NormalHouse) setNumFloor() {
	n.Floor = 2
}

func (n *NormalHouse) getHouse() House {
	return n.House
}
