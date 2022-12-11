package behavioralPatterns

// A type that stores the list of mementos that can have the logic to store them in a database or
// to not store more than a specified number of them
type Caretaker struct {
	mementoArray []*Memento
}

func (c *Caretaker) addMemento(m *Memento) {
	c.mementoArray = append(c.mementoArray, m)
}

func (c *Caretaker) getMemento(index int) *Memento {
	if index >= len(c.mementoArray) || index < 0 {
		return &Memento{}
	}
	return c.mementoArray[index]
}
