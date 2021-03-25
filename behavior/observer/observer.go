package observer

import "fmt"

type EVT interface {
	PrintMesg()
}

type Event struct {
	Message string
}

func (e *Event) PrintMesg() {
	fmt.Printf(e.Message)
}

type Observer interface {
	Notify(event EVT)
}

type Publisher struct {
	ObserversList []Observer
}

func (p *Publisher) AddObserver(o Observer) {
	p.ObserversList = append(p.ObserversList, o)
}

func (p *Publisher) RemoveObserver(o Observer) {
	var indexToRemove int

	for i, observer := range p.ObserversList {
		if observer == o {
			indexToRemove = i
			break
		}
	}

	// slice = append(slice, anotherSlice...)
	p.ObserversList = append(p.ObserversList[:indexToRemove], p.ObserversList[indexToRemove+1:]...)
}

func (s *Publisher) NotifyObservers(event EVT) {
	fmt.Printf("Publisher received message '%s' to notify observers\n", event)
	for _, observer := range s.ObserversList {
		observer.Notify(event)
	}
}
