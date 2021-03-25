package main

import (
	"strconv"
	"strings"
)

const (
	SUM = "sum"
	SUB = "sub"
)

type polishNotationStack []Interpreter

func (s *polishNotationStack) Push(i Interpreter) {
	*s = append(*s, i)
}

func (s *polishNotationStack) Pop() Interpreter {
	length := len(*s)

	if length > 0 {
		temp := (*s)[length-1]
		*s = (*s)[:length-1]
		return temp
	}

	return nil
}

type Interpreter interface {
	Read() int
}

type value int

func (v *value) Read() int {
	return int(*v)
}

type operationSum struct {
	Left  Interpreter
	Right Interpreter
}

func (a *operationSum) Read() int {
	return a.Left.Read() + a.Right.Read()
}

type operationSubtract struct {
	Left  Interpreter
	Right Interpreter
}

func (s *operationSubtract) Read() int {
	return s.Left.Read() - s.Right.Read()
}

func operatorFactory(o string, left, right Interpreter) Interpreter {
	switch o {
	case SUM:
		return &operationSum{
			Left:  left,
			Right: right,
		}
	case SUB:
		return &operationSubtract{
			Left:  left,
			Right: right,
		}
	}

	return nil
}

func main() {
	stack := polishNotationStack{}
	operators := strings.Split("3 4 sum 2 sub", " ")

	for _, operatorString := range operators {
		if operatorString == SUM || operatorString == SUB {
			right := stack.Pop()
			left := stack.Pop()
			mathFunc := operatorFactory(operatorString, left, right)
			res := value(mathFunc.Read())
			stack.Push(&res)
		} else {
			val, err := strconv.Atoi(operatorString)
			if err != nil {
				panic(err)
			}

			temp := value(val)
			stack.Push(&temp)
		}
	}

	println(int(stack.Pop().Read()))
}
