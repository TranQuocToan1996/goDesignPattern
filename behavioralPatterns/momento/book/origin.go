package behavioralPatterns

type originator struct {
	state State
}

func (o *originator) NewMemento() memento {
	return memento{state: o.state}
}
func (o *originator) ExtractAndStoreState(m memento) {
	o.state = m.restore()
}
