package creationalpatterns

import "testing"

func TestBuilderPatterns(t *testing.T) {
	// in charge of the creation of every vehicle during the test
	manufacturingComplex := ManufacturingDirector{}

	carBuilder := &CarBuilder{}
	// CarBuilder that we then passed to manufacturing by using the SetBuilder method
	manufacturingComplex.SetBuilder(carBuilder)
	// we can call the Construct method to create the VehicleProduct using CarBuilder
	manufacturingComplex.Constructs()

	// just return
	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Errorf("Wheels on a car must be 4 and they were %d\n", car.Wheels)
	}
	if car.Structure != "Car" {
		t.Errorf("Structure on a car must be 'Car' and was %s\n",
			car.Structure)
	}
	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5 and they were %d\n", car.Seats)
	}

	bikeBuilder := &BikeBuilder{}
	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Constructs()
	motorbike := bikeBuilder.GetVehicle()
	motorbike.Seats = 1
	if motorbike.Wheels != 2 {
		t.Errorf("Wheels on a motorbike must be 2 and they were %d\n",
			motorbike.Wheels)
	}
	if motorbike.Structure != "Motorbike" {
		t.Errorf("Structure on a motorbike must be 'Motorbike' and was %s\n",
			motorbike.Structure)
	}

}
