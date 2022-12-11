package behavioralPatterns

type Originator struct {
	state string
}

// TODO: The return type of the method should be of the interface
func (e *Originator) createMemento() *Memento {
	return &Memento{state: e.state}
}

// TODO: Should accept interface{}
func (e *Originator) restoreMemento(m *Memento) {
	e.state = m.getSavedState()
}

func (e *Originator) setState(state string) {
	e.state = state
}

func (e *Originator) getState() string {
	return e.state
}
