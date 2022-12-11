package withcommandandfacadepattern

type careTaker struct {
	mementoList []Memento
}

func (c *careTaker) Add(m Memento) {
	c.mementoList = append(c.mementoList, m)
}
func (c *careTaker) Pop() Memento {
	if len(c.mementoList) > 0 {
		tempMemento := c.mementoList[len(c.mementoList)-1]
		c.mementoList = c.mementoList[0 : len(c.mementoList)-1]
		return tempMemento
	}
	return Memento{}
}

func (c *careTaker) Memento(i int) Memento {
	if len(c.mementoList) < i || i < 0 {
		return Memento{}
	}
	return c.mementoList[i]
}
