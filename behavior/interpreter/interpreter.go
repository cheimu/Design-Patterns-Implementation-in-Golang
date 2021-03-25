package interpreter

import (
	"strconv"
	"strings"
)

type polishNotationStack []int

func (s *polishNotationStack) Push(val int) {
	*s = append(*s, val)
}

func (s *polishNotationStack) Pop() int {
	length := len(*s)
	if length > 0 {
		temp := (*s)[length-1]
		(*s) = (*s)[:length-1]
		return temp
	}
	return 0
}

const (
	SUM = "sum"
	SUB = "sub"
	MUL = "mul"
	DIV = "div"
)

func getOperationFunc(o string) func(a, b int) int {
	switch o {
	case SUM:
		return func(a, b int) int {
			return a + b
		}
	case SUB:
		return func(a, b int) int {
			return a - b
		}
	case MUL:
		return func(a, b int) int {
			return a * b
		}
	case DIV:
		return func(a, b int) int {
			return a / b
		}
	}

	return nil
}

func isOperator(o string) bool {
	if o == SUM || o == SUB || o == MUL || o == DIV {
		return true
	}

	return false
}

func Calculate(o string) (int, error) {
	s := polishNotationStack{}
	operators := strings.Split(o, " ")
	for _, operatorString := range operators {
		if isOperator(operatorString) {
			right := s.Pop()
			left := s.Pop()
			mathFunc := getOperationFunc(operatorString)
			res := mathFunc(left, right)
			s.Push(res)
		} else {
			val, err := strconv.Atoi(operatorString)
			if err != nil {
				return 0, err
			}
			s.Push(val)
		}
	}
	return s.Pop(), nil
}
