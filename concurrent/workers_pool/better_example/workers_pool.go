package workers_pool

import (
	"fmt"
	"time"
)

const (
	WORKERS_SIZE              = 10
	TIMEOUT_TIME_MILLISECONDS = 1000
)

//dispatcher will receive messages and use the availableWorkers to get idle workers
func Dispatcher(c *chan int, idleWorkers *chan *chan int) {
	for {
		value := <-*c

		//Check if we don't have idle works so we discard the message
		if len(*idleWorkers) == 0 {
			dropCall(value)
			continue
		}

		worker := <-*idleWorkers
		*worker <- value
	}
}

//createWorkers will create N goroutines to process jobs concurrently
func CreateWorkers(idleWorkers *chan *chan int) {
	for i := 0; i < WORKERS_SIZE; i++ {
		ch := make(chan int)
		go worker(&ch, idleWorkers)
		*idleWorkers <- &ch
	}
}

//worker receives a channel where the jobs will come from. He has to return the
//channel to available workers after use
func worker(c *chan int, idleWorkers *chan *chan int) {
	for {
		value := <-*c
		doActuallyCall(value)
		time.Sleep(time.Millisecond * TIMEOUT_TIME_MILLISECONDS)
		*idleWorkers <- c
	}
}

func doActuallyCall(i int) {
	fmt.Printf("%04d: Called\n", i)
}

func dropCall(i int) {
	fmt.Printf("%04d: Dropped\n", i)
}
