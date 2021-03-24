package memento

import "fmt"

// Originator loads from and stores into Memento
// CareTaker loads and stores a list of Memento

// Memento
type memento struct {
	state State
}

type State struct {
	Description string
}

// Originator
type Originator struct {
	state State
}

func (o *Originator) NewMemento() memento {
	return memento{state: o.state}
}

func (o *Originator) ExtractAndLoadState(m memento) {
	o.state = m.state
}

// CareTaker
type CareTaker struct {
	mementoList []memento
}

func (c *CareTaker) Add(m memento) {
	c.mementoList = append(c.mementoList, m)
}

func (c *CareTaker) Memento(i int) (memento, error) {
	if len(c.mementoList) < i || i < 0 {
		return memento{}, fmt.Errorf("index not found")
	}
	return c.mementoList[i], nil
}
