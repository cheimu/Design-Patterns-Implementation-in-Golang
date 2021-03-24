package main

import "fmt"

// -----------------------------  Object Command ----------------------------------
type Command interface {
	GetValue() interface{}
}

type Volume byte

func (v Volume) GetValue() interface{} {
	return v
}

type Mute bool

func (m Mute) GetValue() interface{} {
	return m
}

// -----------------------------  Memento ----------------------------------
type Memento struct {
	memento Command
}

type Originator struct {
	Command Command
}

func (o *Originator) NewMemento() Memento {
	return Memento{memento: o.Command}
}

func (o *Originator) ExtractAndLoadCommand(m Memento) {
	o.Command = m.memento
}

type CareTaker struct {
	mementoStack []Memento
}

func (c *CareTaker) Push(m Memento) {
	c.mementoStack = append(c.mementoStack, m)
}

func (c *CareTaker) Pop() Memento {
	if len(c.mementoStack) > 0 {
		memento := c.mementoStack[len(c.mementoStack)-1]
		c.mementoStack = c.mementoStack[0 : len(c.mementoStack)-1]
		return memento
	}

	return Memento{}
}

// -----------------------------  Facade ----------------------------------

type MementoFacade struct {
	originator Originator
	careTaker  CareTaker
}

func (m *MementoFacade) SaveSettings(s Command) {
	m.originator.Command = s
	m.careTaker.Push(m.originator.NewMemento())
}

func (m *MementoFacade) RestoreSettings() Command {
	m.originator.ExtractAndLoadCommand(m.careTaker.Pop())
	return m.originator.Command
}

// -----------------------------  Main ----------------------------------
func main() {
	m := MementoFacade{}

	m.SaveSettings(Volume(4))
	m.SaveSettings(Mute(false))

	assertAndPrint(m.RestoreSettings())
	assertAndPrint(m.RestoreSettings())
}

func assertAndPrint(c Command) {
	switch cast := c.(type) {
	case Volume:
		fmt.Printf("Volume:\t%d\n", cast)
	case Mute:
		fmt.Printf("Mute:\t%t\n", cast)
	}
}
