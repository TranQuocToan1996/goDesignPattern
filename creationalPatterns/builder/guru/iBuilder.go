package creationalpatterns

// https://refactoring.guru/design-patterns/builder

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

func getBuilder(builderType builderType) IBuilder {
	if builderType == normal {
		return newNormalBuilder()
	}

	if builderType == igloo {
		return newIglooBuilder()
	}

	return nil
}
