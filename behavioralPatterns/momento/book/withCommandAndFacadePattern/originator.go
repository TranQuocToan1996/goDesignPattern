package withcommandandfacadepattern

type originator struct {
	Command Command
}

func (o *originator) NewMemento() Memento {
	return Memento{memento: o.Command}
}
func (o *originator) ExtractAndStoreCommand(m Memento) {
	o.Command = m.memento
}
