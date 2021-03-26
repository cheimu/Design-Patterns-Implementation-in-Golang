package worker_pool

import (
	"fmt"
	"strings"
)

// WorkLauncher is the type that can handle Request types
type WorkLauncher interface {
	LaunchWorker(id int, in chan Request)
}

type PreffixSuffixWorker struct {
	ID      int
	PrefixS string
	SuffixS string
}

func (w *PreffixSuffixWorker) LaunchWorker(id int, in chan Request) {
	fmt.Printf("Now Launch Work %d\n", w.ID)
	w.prefix(w.append(w.uppercase(in)))
}

func (w *PreffixSuffixWorker) uppercase(in <-chan Request) <-chan Request {
	out := make(chan Request)
	go func() {
		for msg := range in {
			s, ok := msg.Data.(string)
			if !ok {
				msg.Handler(nil)
				continue
			}
			msg.Data = strings.ToUpper(s)
			out <- msg
		}
		close(out)
	}()
	return out
}

func (w *PreffixSuffixWorker) append(in <-chan Request) <-chan Request {
	out := make(chan Request)
	go func() {
		for msg := range in {
			uppercaseString, ok := msg.Data.(string)
			if !ok {
				msg.Handler(nil)
				continue
			}
			msg.Data = fmt.Sprintf("%s%s", uppercaseString, w.SuffixS)
			out <- msg
		}
		close(out)
	}()
	return out
}

func (w *PreffixSuffixWorker) prefix(in <-chan Request) {
	go func() {
		for msg := range in {
			uppercasedStringWithSuffix, ok := msg.Data.(string)
			if !ok {
				msg.Handler(nil)
				continue
			}
			msg.Handler(fmt.Sprintf("%s%s", w.PrefixS, uppercasedStringWithSuffix))
		}
	}()
}
