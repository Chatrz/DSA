package linkedlist

import "testing"

func testNewLinkedList(T *testing.T) {
	l := NewLinkedList()
	l.Append(2)
	l.Prepend(3)
	if val, err := l.Get(0); err != nil || val.Value != 3 {
		T.Errorf("riiidi")
	}
}
