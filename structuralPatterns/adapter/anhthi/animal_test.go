package anhthi

import "testing"

func TestAnimal(t *testing.T) {
	var animals []Animal
	animals = append(animals, new(Cat))
	animals = append(animals, NewAdapterCrocodile())

	for _, animal := range animals {
		animal.Move()
	}
}
