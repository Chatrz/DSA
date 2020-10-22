package stack

import "errors"

const LIMIT = 1000

type Stack struct {
	stack []int
	top   int
}

func NewStack() *Stack {
	s := new(Stack)
	s.top = -1
	return s
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

func (s *Stack) Push(key int) error {
	if s.Size() > LIMIT {
		return errors.New("stack overflow error")
	}
	s.top++
	if s.top == len(s.stack) {
		s.stack = append(s.stack, key)
	} else {
		s.stack[s.top] = key
	}
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return -1, errors.New("stack underflow error")
	}
	s.top--
	return s.stack[s.top+1], nil
}

func (s *Stack) Peek() int {
	return s.stack[s.top]
}

func (s *Stack) Size() int {
	return len(s.stack)
}
