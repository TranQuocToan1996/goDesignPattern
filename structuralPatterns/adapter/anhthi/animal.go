package anhthi

type Animal interface {
	Move()
}

// Cat is a concrete animal since it implements the method Move
type Cat struct{}

func (c *Cat) Move() {}

// and somewhere in the code we need to use the crocodile type which is often not our code and this Crocodile type does not implement the Animal interface
// but we need to use a crocodile as an animal

type Crocodile struct{}

func (c *Crocodile) Slither() {}

// we create a CrocodileAdapter struct that adapts a composite crocodile so that it can be used as an Animal

type CrocodileAdapter struct {
	c *Crocodile
}

func NewAdapterCrocodile() *CrocodileAdapter {
	return &CrocodileAdapter{new(Crocodile)}
}

func (adapter *CrocodileAdapter) Move() {
	adapter.c.Slither()
}
