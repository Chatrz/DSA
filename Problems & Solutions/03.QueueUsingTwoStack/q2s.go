/*
 * +------------------------------------+
 * | Author:        hadi abbasi	        |
 * |                                    |
 * | Link:          github.com/hawwwdi  |
 * +------------------------------------+
 */

package q2s

import stack "github.com/sinamna/DSA/Structures/Stack/Go/hadi"

type Q2s struct {
	pusher *stack.Stack
	poper  *stack.Stack
}

func NewQ2s() *Q2s {
	return &Q2s{
		pusher: stack.NewStack(),
		poper:  stack.NewStack(),
	}
}

func (q *Q2s) Enqueue(key int) error {
	for k, err := q.poper.Pop(); err == nil; k, err = q.poper.Pop() {
		if err1 := q.pusher.Push(k); err1 != nil {
			return err1
		}
	}
	if err := q.pusher.Push(key); err != nil {
		return err
	}
	for k, err := q.pusher.Pop(); err == nil; k, err = q.pusher.Pop() {
		if err1 := q.poper.Push(k); err != nil {
			return err1
		}
	}
	return nil
}

func (q *Q2s) Dequeue(key int) (int, error) {
	return q.poper.Pop()
}
