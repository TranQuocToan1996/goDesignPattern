package creationalpatterns

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

type IglooBuilder struct {
	windowType windowType
	doorType   doorType
	floor      int
}

func (n *IglooBuilder) setWindowType() {
	n.windowType = "Snow Window"
}

func (n *IglooBuilder) setDoorType() {
	n.doorType = "Snow Door"
}

func (n *IglooBuilder) setNumFloor() {
	n.floor = 1
}

func (n *IglooBuilder) getHouse() House {
	return House{
		windowType: n.windowType,
		doorType:   n.doorType,
		floor:      n.floor,
	}
}
