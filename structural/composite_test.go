package composite

import "testing"

func TestAthlete_Train(t *testing.T) {
	athlete := Athlete{}
	athlete.Train()
}

func TestSwimmer_Swim(t *testing.T) {
	swim := Swim
	swimmer := CompositeSwimmerA{
		MySwim: &swim,
	}
	swimmer.MyAthlete.Train()
	(*swimmer.MySwim)()
}

func TestSwimmer_Swim2(t *testing.T) {
	swimmer := CompositeSwimmerB{
		&Athlete{},
		&SwimmerImplementor{},
	}

	swimmer.Train()
	swimmer.Swim()
}

func TestGetParent(t *testing.T) {
	son := &Son{}
	son.GetParent()
}
