package linkedlist

import "errors"

type Node struct {
	Value int
	next  *Node
	prev  *Node
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

//LinkedList linkedlist
type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func (l *LinkedList) Size() int {
	return l.size
}

func NewLinkedList() *LinkedList {
	l := new(LinkedList)
	return l
}

func (l *LinkedList) Append(value int) {
	l.size++
	n := new(Node)
	n.Value = value
	n.prev = l.tail
	l.tail = n
	if l.head == nil {
		l.head = n
	}
}

func (l *LinkedList) Prepend(value int) {
	l.size++
	n := new(Node)
	n.Value = value
	if l.head == nil {
		l.head = n
		l.tail = n
		return
	}
	n.next = l.head
	l.head.prev = n
	l.head = n
}

func (l *LinkedList) Get(index int) (*Node, error) {
	if index > l.size {
		return nil, errors.New("invalid index")
	}
	n := l.head
	for i := 0; i < index; i++ {
		n = n.Next()
	}
	return n, nil
}

func (l *LinkedList) Remove(index int) bool {
	if index == 0 {
		l.head = l.head.Next()
		l.head.prev = nil
		l.size--
		return true
	}
	if index == l.size {
		l.tail = l.tail.prev
		l.tail.next = nil
		l.size--
		return true
	}
	n, err := l.Get(index)
	if err != nil {
		return false
	}
	n.next.prev = n.Prev()
	n.prev.next = n.Next()
	l.size--
	return true
}

//todo add Insert, Head, Tail method
