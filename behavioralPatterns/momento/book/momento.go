package behavioralPatterns

// don't store directly State instances. Basically, because it will couple the originator and the careTaker to the business object and we want to have as little coupling as possible
type memento struct {
	state State
}

func (m memento) restore() State {
	return m.state
}
