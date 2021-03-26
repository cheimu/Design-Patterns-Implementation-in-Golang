package main

import (
	"Go-Design-Patterns/concurrent/worker_pool"
	"fmt"
	"sync"
)

func main() {
	bufferSize := 100
	var dispatcher worker_pool.Dispatch = worker_pool.NewDispatcher(bufferSize)
	workers := 3
	for i := 0; i < workers; i++ {
		worker := worker_pool.PreffixSuffixWorker{
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
		req := worker_pool.NewStringRequest(mesg, i, &wg)
		dispatcher.MakeRequest(req)
	}
	dispatcher.Stop()
	wg.Wait()

}
