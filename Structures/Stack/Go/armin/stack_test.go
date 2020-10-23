package stack

import (
	"fmt"
	"testing"
)

func TestCreateStack(t *testing.T) {
	stack:=CreateStack()
	if stack.top!=-1{
		t.Error("CreateStack does not work !")
	}
}

func TestStack_Push(t *testing.T) {
	stack:=CreateStack()
	stack.Push(20)
	stack.Push(30)
	if stack.stack[1]!=30{
		t.Error("Push does not work !")
	}
}

func TestStack_Pop(t *testing.T) {
	stack:=CreateStack()
	stack.Push(20)
	stack.Push(30)
	stack.Push(30)
	stack.Push(30)
	stack.Push(30)
	stack.Pop()
	stack.Pop()
	fmt.Println(stack.stack)
	if len(stack.stack)!=3{
		t.Error("Pop does not work !")
	}
}
func TestStack_Peak(t *testing.T) {
	stack:=CreateStack()
	stack.Push(20)
	stack.Push(30)
	if stack.Peak()!=30{
		t.Error("Peak does not work !")
	}
}