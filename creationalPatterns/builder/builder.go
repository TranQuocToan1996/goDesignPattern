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

type ManufacturingDirector struct {
	builder BuildProcess
}

// Construct method that will use the builder that is stored in Manufacturing, and will reproduce the required steps
func (m *ManufacturingDirector) Constructs() {
	m.builder.SetSeats().SetStructure().SetWheels()
}

// The SetBuilder method will allow us to change the builder that is being used in the Manufacturing director
func (m *ManufacturingDirector) SetBuilder(b BuildProcess) {
	m.builder = b
}

type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}
func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}
func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

// The process are difference. Diff logic
type BikeBuilder struct {
	v VehicleProduct
}

func (c *BikeBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 2
	return c
}

func (c *BikeBuilder) SetSeats() BuildProcess {
	c.v.Seats = 2
	return c
}

func (c *BikeBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Motorbike"
	return c
}

func (c *BikeBuilder) GetVehicle() VehicleProduct {
	return c.v
}

type MotorBuilder struct {
	v VehicleProduct
}

func (m *MotorBuilder) SetWheels() BuildProcess {
	m.v.Wheels = 8
	return m
}

func (m *MotorBuilder) SetSeats() BuildProcess {
	m.v.Seats = 30
	return m
}

func (m *MotorBuilder) SetStructure() BuildProcess {
	m.v.Structure = "Bus"
	return m
}

func (m *MotorBuilder) GetVehicle() VehicleProduct {
	return m.v
}
