/*
 * +------------------------------------+
 * | Author:        hadi abbasi	        |
 * |                                    |
 * | Link:          github.com/hawwwdi  |
 * +------------------------------------+
 */

package queue

import "errors"

type Queue struct {
	queue    []int
	capacity int
	size     int
	head     int
	tail     int
}

func NewQueue(capacity int) *Queue {
	q := new(Queue)
	q.queue = make([]int, capacity)
	q.capacity = capacity
	q.tail = capacity - 1
	return q
}

func (q *Queue) Enqueue(key int) error {
	if q.IsFull() {
		return errors.New("queue overflow error")
	}
	q.tail = (q.tail + 1) % q.capacity
	q.queue[q.tail] = key
	q.size++
	return nil
}

func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return -1, errors.New("queue underfolow error")
	}
	value := q.queue[q.head]
	q.head = (q.head + 1) % q.capacity
	q.size--
	return value, nil
}

func (q *Queue) Head() (int, error) {
	if q.IsEmpty() {
		return -1, errors.New("queue underflow error")
	}
	return q.queue[q.head], nil
}

func (q *Queue) Tail() (int, error) {
	if q.IsEmpty() {
		return -1, errors.New("queue underflow error")
	}
	return q.queue[q.tail], nil
}

func (q *Queue) IsFull() bool {
	return q.size == q.capacity
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Size() int {
	return q.size
}
