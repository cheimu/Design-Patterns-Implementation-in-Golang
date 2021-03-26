package concurrent_observer

import (
	"fmt"
	"io"
	"os"
	"time"
)

// ---------------------------- Subscriber ---------------------------------
type Subscriber interface {
	Notify(interface{}) error
	Close()
}

type writerSubscriber struct {
	in     chan interface{}
	id     int
	Writer io.Writer
}

func NewWriterSubscriber(id int, out io.Writer) Subscriber {
	if out == nil {
		out = os.Stdout
	}

	s := &writerSubscriber{
		in:     make(chan interface{}),
		id:     id,
		Writer: out,
	}

	go func() {
		for msg := range s.in {
			fmt.Fprintf(s.Writer, "(W%d): %v\n", s.id, msg)
		}
	}()
	return s
}

func (s *writerSubscriber) Notify(msg interface{}) (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%#v", rec)
		}
	}()
	select {
	case s.in <- msg:
	case <-time.After(time.Second):
		err = fmt.Errorf("timeout")
	}
	return
}

func (s *writerSubscriber) Close() {
	close(s.in)
}

// ---------------------------- Publisher ---------------------------------
type Publisher interface {
	start()
	AddSubscriberCh() chan Subscriber
	RemoveSubscriberCh() chan Subscriber
	PublishingCh() chan interface{}
	Stop() chan interface{}
}

func NewPublisher() Publisher {
	p := &publisher{
		subscribers: make([]Subscriber, 0),
		addSubCh:    make(chan Subscriber),
		removeSubCh: make(chan Subscriber),
		in:          make(chan interface{}),
		stop:        make(chan interface{}),
	}
	go p.start()
	return p
}

type publisher struct {
	subscribers []Subscriber
	addSubCh    chan Subscriber
	removeSubCh chan Subscriber
	in          chan interface{}
	stop        chan interface{}
}

func (p *publisher) AddSubscriberCh() chan Subscriber {
	return p.addSubCh
}
func (p *publisher) RemoveSubscriberCh() chan Subscriber {
	return p.removeSubCh
}
func (p *publisher) PublishingCh() chan interface{} {
	return p.in
}

func (p *publisher) start() {
	for {
		select {
		case msg := <-p.in:
			for _, sub := range p.subscribers {
				sub.Notify(msg)
			}
		case sub := <-p.addSubCh:
			p.subscribers = append(p.subscribers, sub)
		case sub := <-p.removeSubCh:
			for i, candidate := range p.subscribers {
				if candidate == sub {
					p.subscribers = append(p.subscribers[:i],
						p.subscribers[i+1:]...)
					candidate.Close()
					break
				}
			}
		case <-p.stop:
			for _, sub := range p.subscribers {
				sub.Close()
			}
			close(p.addSubCh)
			close(p.in)
			close(p.removeSubCh)
			return
		}
	}
}

func (p *publisher) Stop() chan interface{} {
	return p.stop
}
