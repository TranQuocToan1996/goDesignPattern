package behavioralPatterns

import "testing"

const (
	idle = "Idle"
)

func TestCareTaker_Add(t *testing.T) {
	originator := originator{}

	originator.state = State{Description: idle}
	careTaker := careTaker{}
	mem := originator.NewMemento()
	if mem.state.Description != idle {
		t.Error("Expected state was not found")
	}
	currentLen := len(careTaker.mementoList)
	careTaker.Add(mem)
	if len(careTaker.mementoList) != currentLen+1 {
		t.Error("No new elements were added on the list")
	}
}

func TestCareTaker_Memento(t *testing.T) {
	originator := originator{}
	careTaker := careTaker{}
	originator.state = State{idle}
	careTaker.Add(originator.NewMemento())
	mem, err := careTaker.Memento(0)
	if err != nil {
		t.Fatal(err)
	}
	if mem.state.Description != idle {
		t.Error("Unexpected state")
	}

	mem, err = careTaker.Memento(-1)
	if err == nil {
		t.Fatal("An error is expected when asking for a negative number but no error was found")
	}
}

func TestOriginator_ExtractAndStoreState(t *testing.T) {
	originator := originator{state: State{idle}}
	idleMemento := originator.NewMemento()
	originator.ExtractAndStoreState(idleMemento)
	if originator.state.Description != idle {
		t.Error("Unexpected state found")
	}
}
