package stack

import "fmt"

type Stack struct {
	stack []int
	top   int
}

func CreateStack() *Stack {
	stack := &Stack{
		stack: []int{},
		top:   -1,
	}
	return stack
}

func (stack *Stack) IsEmpty() bool {
	if stack.top == -1 {
		fmt.Println("STACK IS EMPTY !")
		return true
	}
	return false
}

func (stack *Stack) Push(element int) {
	stack.stack = append(stack.stack, element)
	stack.top += 1
}

func (stack *Stack) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	} else {
		popped := stack.stack[stack.top]
		stack.top -= 1
		stack.stack = stack.stack[:stack.top+1]
		return popped
	}
}

func (stack *Stack) Peak() interface{} {
	if stack.IsEmpty() {
		return nil
	} else {
		return stack.stack[stack.top]
	}
}
