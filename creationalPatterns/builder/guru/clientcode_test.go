package creationalpatterns

import (
	"testing"
)

func TestGuruBuilder(t *testing.T) {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(normalBuilder)
	normalHouse := director.build()

	if normalHouse.doorType != WoodenDoor || normalHouse.floor != 2 ||
		normalHouse.windowType != WoodenWindow {
		t.Error("wrong builder")
	}

	// fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	// fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	// fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.build()

	if iglooHouse.doorType != SnowDoor || iglooHouse.floor != 1 ||
		iglooHouse.windowType != SnowWindow {
		t.Error("wrong builder")
	}

	// fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
	// fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
	// fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor)

}
