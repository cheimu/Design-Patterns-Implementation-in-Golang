package mediator

import (
	"testing"
)

func TestMediator(t *testing.T) {

	mediator := NewMediator()
	mediator.Ted.Talk()

	if "Ted: Bill?\n"+"Bill: What?\n"+"Ted: Strange things are afoot at the Circle K.\n" != mediator.Received {
		t.Errorf(mediator.Received)
	}

}
