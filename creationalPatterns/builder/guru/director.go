package creationalpatterns

func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

type Director struct {
	builder IBuilder
}

func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

// Contructor func
func (d *Director) build() House {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()

	return d.builder.getHouse()
}
