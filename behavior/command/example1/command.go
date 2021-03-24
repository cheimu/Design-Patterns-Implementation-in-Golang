package main

import (
	"fmt"
)

type Command interface {
	Execute()
}

type ConsoleOutput struct {
	message string
}

func (c *ConsoleOutput) Execute() {
	fmt.Println(c.message)
}

func CreateCommand(s string) Command {
	fmt.Println("Creating command")

	return &ConsoleOutput{
		message: s,
	}
}

type CommandQueue struct {
	queue []Command
}

func (q *CommandQueue) AddCommand(c Command) {
	q.queue = append(q.queue, c)
	if len(q.queue) == 3 {
		q.queue[0].Execute()
		q.queue = q.queue[1:]
	}
}

func main() {
	queue := CommandQueue{}

	queue.AddCommand(CreateCommand("First message"))
	queue.AddCommand(CreateCommand("Second message"))
	queue.AddCommand(CreateCommand("Third message"))

	queue.AddCommand(CreateCommand("Fourth message"))
	queue.AddCommand(CreateCommand("Fifth message"))

}
