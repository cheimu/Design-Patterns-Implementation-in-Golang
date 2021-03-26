package workers_pool

import (
	"fmt"
	"log"
	"sync"
)

// Request is the type of workload that workers need to work with
type Request struct {
	Data    interface{}
	Handler RequestHandler
}

type RequestHandler func(interface{})

func NewStringRequest(s string, id int, wg *sync.WaitGroup) Request {
	myRequest := Request{
		Data: s,
		Handler: func(i interface{}) {
			defer wg.Done()
			s, ok := i.(string)
			if !ok {
				log.Fatal("Invalid casting to string")
			}
			fmt.Println(s)
		},
	}

	return myRequest
}
