package structuralPatterns

import "fmt"

type Athlete struct{}

func (a *Athlete) Train() {
	fmt.Println("Training")
}

// Directly
type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    func()
}

func Swim() {
	fmt.Println("Swimming!")
}

/* localSwim := Swim
   swimmer := CompositeSwimmerA{
     MySwim: localSwim,
}
   swimmer.MyAthlete.Train()
   swimmer.MySwim () */

type Animal struct{}

func (r *Animal) Eat() {
	println("Eating")
}

// Embed
type Shark struct {
	Animal
	Swim func()
}

/*   fish := Shark{
     Swim: Swim,
}
   fish.Eat()
   fish.Swim() */

type Swimmer interface {
	Swim()
}
type Trainer interface {
	Train()
}
type SwimmerImpl struct{}

func (s *SwimmerImpl) Swim() {
	println("Swimming!")
}

type CompositeSwimmerB struct {
	Trainer
	Swimmer
}

// swimmer := CompositeSwimmerB{
//      &Athlete{},
//      &SwimmerImpl{},
// }
//    swimmer.Train()
//    swimmer.Swim()
