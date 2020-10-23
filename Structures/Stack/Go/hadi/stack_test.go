package stack

import "testing"

func TestStack(t *testing.T) {
	tests := []struct {
		keys []int
		pops []int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{5, 4, 3, 2, 1},
		},
	}
	for _, test := range tests {
		stack := NewStack()
		for _, key := range test.keys {
			stack.Push(key)
		}
		for _, key := range test.pops {
			peeked := stack.Peek()
			poped, _ := stack.Pop()
			if peeked != key || poped != key {
				t.Errorf("%v != %v", peeked, poped)
			}
		}
	}
}

func TestSize(t *testing.T) {
	tests := []struct {
		keys []int
		size int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			5,
		},
		{
			[]int{},
			0,
		},
	}
	for _, test := range tests {
		stack := NewStack()
		for _, keys := range test.keys {
			stack.Push(keys)
		}
		if stack.Size() != test.size {
			t.Errorf("%v != %v", stack.Size(), test.size)
		}
	}
}

func TestOverFlow(t *testing.T) {
	stack := NewStack()
	for i := 0; i < 1000; i++ {
		err := stack.Push(i)
		if err != nil {
			t.Error("there is a bug")
		}
	}
	err := stack.Push(limit + 1)
	if err == nil {
		t.Errorf("there is a bug stack size is %v", stack.Size())
	}
}
