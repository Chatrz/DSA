/*
 * +------------------------------------+
 * | Author:        hadi abbasi			|
 * |									|
 * | Link:          github.com/hawwwdi  |
 * +------------------------------------+
 */

package main

import (
	"fmt"

	queue "github.com/sinamna/DSA_Trainings/Structures/Queue/Go"
)

func main() {
	q := queue.NewQueue(5)
	for i := 1; i < 6; i++ {
		q.Enqueue(i)
	}
	res := reverse(q)
	for x, err := res.Dequeue(); err == nil; x, err = res.Dequeue() {
		fmt.Printf("%v, ", x)
	}
}

func reverse(q *queue.Queue) *queue.Queue {
	ans := queue.NewQueue(q.Size())
	for q.Size() != 0 {
		for j := 0; j < q.Size()-1; j++ {
			x, _ := q.Dequeue()
			q.Enqueue(x)
		}
		x, _ := q.Dequeue()
		ans.Enqueue(x)
	}
	return ans
}
