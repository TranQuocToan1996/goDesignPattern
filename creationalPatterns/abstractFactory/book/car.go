package creationalPatterns

import (
	"fmt"
)

type Car interface {
	NumDoors() int
}

const (
	LuxuryCarType = 1
	FamilyCarType = 2
)

type CarFactory struct{}

func (c *CarFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, fmt.Errorf("vehicle of type %d not recognized", v)
	}
}

type LuxuryCar struct{}

func (*LuxuryCar) NumDoors() int {
	return 4
}
func (*LuxuryCar) NumWheels() int {
	return 4
}
func (*LuxuryCar) NumSeats() int {
	return 5
}

type FamilyCar struct{}

func (*FamilyCar) NumDoors() int {
	return 5
}
func (*FamilyCar) NumWheels() int {
	return 4
}
func (*FamilyCar) NumSeats() int {
	return 5
}
