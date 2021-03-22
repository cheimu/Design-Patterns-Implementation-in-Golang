package composite

import "fmt"

//-------------------- Example 1 --------------------------------------
type Athlete struct{}

func (a *Athlete) Train() {
	fmt.Println("Training")
}

func Swim() {
	fmt.Println("Swiming")
}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    *func()
}

//-------------------- Example 2 --------------------------------------

type Trainer interface {
	Train()
}

type Swimmer interface {
	Swim()
}

type SwimmerImplementor struct{}

func (s *SwimmerImplementor) Swim() {
	println("Swimming!")
}

type CompositeSwimmerB struct {
	Trainer
	Swimmer
}

//-------------------- Example 3 --------------------------------------
type Parent struct {
	FieldA int
}

type Son struct {
	P Parent
}

func (s *Son) GetParent() Parent {
	return s.P
}
