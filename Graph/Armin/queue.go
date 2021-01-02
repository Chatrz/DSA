package main

import (
	"errors"
)

// customized implemention of queue for bfs algorithm :
type Queue struct {
	queue    []*Vertex
	capacity int
	size     int
	head     int
	tail     int
}

func NewQueue(capacity int) *Queue {
	q := new(Queue)
	q.queue = make([]*Vertex, capacity)
	q.capacity = capacity
	q.tail = capacity - 1
	return q
}

func (q *Queue) Enqueue(key *Vertex) error {
	if q.IsFull() {
		return errors.New("queue overflow error")
	}
	q.tail = (q.tail + 1) % q.capacity
	q.queue[q.tail] = key
	q.size++
	return nil
}

func (q *Queue) Dequeue() (*Vertex, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue underfolow error")
	}
	value := q.queue[q.head]
	q.head = (q.head + 1) % q.capacity
	q.size--
	return value, nil
}

func (q *Queue) IsFull() bool {
	return q.size == q.capacity
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

// main implementation is from my dear friend hadi abbasi ==>  github.com/hawwwdi
////////////////////////////////////////////////////////////////////
