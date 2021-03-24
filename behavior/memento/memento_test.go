package memento

import "testing"

func TestCareTaker_Add(t *testing.T) {
	originator := Originator{}
	originator.state = State{Description: "Idle"}

	careTaker := CareTaker{}

	mem := originator.NewMemento()
	if mem.state.Description != "Idle" {
		t.Error("Expected state was not found")
	}

	currentLen := len(careTaker.mementoList)
	careTaker.Add(mem)

	if len(careTaker.mementoList) != currentLen+1 {
		t.Error("No new elements were added on the list")
	}
}

func TestCareTaker_Memento(t *testing.T) {
	originator := Originator{}
	careTaker := CareTaker{}

	originator.state = State{"Idle"}
	careTaker.Add(originator.NewMemento())

	mem, err := careTaker.Memento(0)
	if err != nil {
		t.Fatal(err)
	}

	if mem.state.Description != "Idle" {
		t.Error("Unexpected state")
	}

	mem, err = careTaker.Memento(-1)
	if err == nil {
		t.Fatal("An error is expected when asking for a negative number but no error was found")
	}
}

func TestOriginator_ExtractAndLoadState(t *testing.T) {
	originator := Originator{state: State{"Idle"}}
	idleMemento := originator.NewMemento()

	originator.state = State{"Working"}

	originator.ExtractAndLoadState(idleMemento)
	if originator.state.Description != "Idle" {
		t.Error("Unexpected state found")
	}
}
