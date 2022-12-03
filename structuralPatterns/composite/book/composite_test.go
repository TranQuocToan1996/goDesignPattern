package structuralPatterns

import "testing"

func TestComposite(t *testing.T) {
	localSwim := Swim
	swimmer := CompositeSwimmerA{
		MySwim: localSwim,
	}
	swimmer.MyAthlete.Train()
	swimmer.MySwim()

	fish := Shark{
		Swim: Swim,
	}
	fish.Eat()
	fish.Swim()

	swimmerB := CompositeSwimmerB{
		&Athlete{},
		&SwimmerImpl{},
	}
	swimmerB.Train()
	swimmerB.Swim()
}
