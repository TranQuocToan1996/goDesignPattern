package creationalpatterns

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

type NormalBuilder struct {
	windowType windowType
	doorType   doorType
	floor      int
}

func (n *NormalBuilder) setWindowType() {
	n.windowType = "Wooden Window"
}

func (n *NormalBuilder) setDoorType() {
	n.doorType = "Wooden Door"
}

func (n *NormalBuilder) setNumFloor() {
	n.floor = 2
}

func (n *NormalBuilder) getHouse() House {
	return House{
		windowType: n.windowType,
		doorType:   n.doorType,
		floor:      n.floor,
	}
}
