package workers_pool

import "time"

// Dispatch is the type that launchs workers in parallel
// and handle all the possible incoming channels
type Dispatch interface {
	LaunchWorker(id int, w WorkLauncher)
	MakeRequest(r Request)
	Stop()
}

// dispatch is an implementation of Dispatch interface
type Dispatcher struct {
	inChan chan Request
}

func (d *Dispatcher) LaunchWorker(id int, w WorkLauncher) {
	w.LaunchWorker(id, d.inChan)
}

func (d *Dispatcher) MakeRequest(r Request) {
	select {
	case d.inChan <- r:
	case <-time.After(time.Second * 5):
		return
	}
}

func (d *Dispatcher) Stop() {
	close(d.inChan)
}

func NewDispatcher(b int) Dispatch {
	return &Dispatcher{
		inChan: make(chan Request, b),
	}
}
