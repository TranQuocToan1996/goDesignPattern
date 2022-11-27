package creationalpatterns

type builderType string

const (
	normal builderType = "normal"
	igloo  builderType = "igloo"
)

type windowType string

const (
	WoodenWindow windowType = "Wooden Window"
	SnowWindow   windowType = "Snow Window"
)

type doorType string

const (
	WoodenDoor doorType = "Wooden Door"
	SnowDoor   doorType = "Snow Door"
)
