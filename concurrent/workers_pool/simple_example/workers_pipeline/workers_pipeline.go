package main

import (
	workers_pool "Go-Design-Patterns/concurrent/workers_pool/simple_example"
	"fmt"
	"sync"
)

func main() {
	bufferSize := 100
	var dispatcher workers_pool.Dispatch = workers_pool.NewDispatcher(bufferSize)
	workers := 3
	for i := 0; i < workers; i++ {
		worker := workers_pool.PreffixSuffixWorker{
			PrefixS: fmt.Sprintf("WorkerID: %d -> ", i),
			SuffixS: " World",
			ID:      i,
		}
		dispatcher.LaunchWorker(i, &worker)
	}
	requests := 10
	var wg sync.WaitGroup
	wg.Add(requests)
	for i := 0; i < requests; i++ {
		mesg := fmt.Sprintf("(Msg_id: %d) -> Hello", i)
		req := workers_pool.NewStringRequest(mesg, i, &wg)
		dispatcher.MakeRequest(req)
	}
	dispatcher.Stop()
	wg.Wait()

}
