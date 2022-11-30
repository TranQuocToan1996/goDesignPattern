package creationalPatterns

import (
	"fmt"
)

type Motorbike interface {
	GetMotorbikeType() int
}

const (
	SportMotorbikeType  = 1
	CruiseMotorbikeType = 2
)

type MotorbikeFactory struct{}

func (c *MotorbikeFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return new(SportMotorbike), nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil
	default:
		return nil, fmt.Errorf("vehicle of type %d not recognized", v)
	}
}

type SportMotorbike struct{}

func (s *SportMotorbike) NumWheels() int {
	return 2
}
func (s *SportMotorbike) NumSeats() int {
	return 1
}
func (s *SportMotorbike) GetMotorbikeType() int {
	return SportMotorbikeType
}

type CruiseMotorbike struct{}

func (c *CruiseMotorbike) NumWheels() int {
	return 2
}
func (c *CruiseMotorbike) NumSeats() int {
	return 2
}
func (c *CruiseMotorbike) GetMotorbikeType() int {
	return CruiseMotorbikeType
}
