package main

import (
	"fmt"
	"time"
)

type Command interface {
	Info() string
}

type TimePassed struct {
	start time.Time
}

func (t *TimePassed) Info() string {
	return time.Since(t.start).String()
}

type HelloMessage struct{}

func (h HelloMessage) Info() string {
	return "Hello world!"
}

func main() {
	timeCommand := &TimePassed{time.Now()}
	helloCommand := &HelloMessage{}

	time.Sleep(time.Second)

	fmt.Println(timeCommand.Info())
	fmt.Println(helloCommand.Info())
}
