package main

import (
	workers_pool "Go-Design-Patterns/concurrent/workers_pool/better_example"
	"math/rand"
	"time"
)

//Lets create an infinite loop of requests every 20 milliseconds so we force the
//discard of rcal
func main() {
	//Concurrency party starts here with initialization. We could use the "init"
	//function but we won't be able to unit test it correctly

	//idleWorkers is an queue of goroutines workers
	idleWorkers := make(chan *chan int, workers_pool.WORKERS_SIZE)

	//A channel to communicate to the job dispatcher
	jobsCh := make(chan int)

	//Create the SPF dispatcher in a different process
	go workers_pool.Dispatcher(&jobsCh, &idleWorkers)

	workers_pool.CreateWorkers(&idleWorkers)

	iter := 0
	for {
		rand.Seed(int64(iter))
		wait := rand.Int31n(200)
		jobsCh <- iter
		iter++
		time.Sleep(time.Duration(wait) * time.Millisecond)
	}
}
