package creationalpatterns

// My notes:
// A Builder design pattern tries to:
// Abstract complex creations so that object creation is separated from the object user
// Create an object step by step by filling its fields and creating the embedded objects
// Reuse the object creation algorithm between many objects

// So the requirements for a Vehicle builder example would be the following:
// I must have a manufacturing type that constructs everything that a vehicle needs When using a car builder, the VehicleProduct with four wheels, five seats, and a structure defined as Car must be returned
// When using a motorbike builder, the VehicleProduct with two wheels, two seats, and a structure defined as Motorbike must be returned
// A VehicleProduct built by any BuildProcess builder must be open to modifications

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

type ManufacturingDirector struct{}

// Construct method that will use the builder that is stored in Manufacturing, and will reproduce the required steps
func (m *ManufacturingDirector) Constructs() {

}

// The SetBuilder method will allow us to change the builder that is being used in the Manufacturing director
func (m *ManufacturingDirector) SetBuilder(b BuildProcess) {

}


type CarBuilder struct{}

func (c *CarBuilder) SetWheels() BuildProcess {
    return nil
}
func (c *CarBuilder) SetSeats() BuildProcess {
    return nil
}
func (c *CarBuilder) SetStructure() BuildProcess {
    return nil
}
func (c *CarBuilder) GetVehicle() VehicleProduct {
    return VehicleProduct{}
}

// The process are difference. Diff logic
type BikeBuilder struct{}

func (c *BikeBuilder) SetWheels() BuildProcess {
    return nil
}
func (c *BikeBuilder) SetSeats() BuildProcess {
    return nil
}
func (c *BikeBuilder) SetStructure() BuildProcess {
    return nil
}
func (c *BikeBuilder) GetVehicle() VehicleProduct {
    return VehicleProduct{}
}