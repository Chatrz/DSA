/*
 * +------------------------------------+
 * | Author:        hadi abbasi	        |
 * |                                    |
 * | Link:          github.com/hawwwdi  |
 * +------------------------------------+
 */

package priorityqueue

type Node struct {
	Priority int
	Data     interface{}
}

type PQ struct {
	heap []Node
}

func NewPQ() *PQ {
	return &PQ{
		heap: make([]Node, 1),
	}
}

func (p *PQ) Enqueue(node Node) {
	p.heap = append(p.heap, node)
	lastChild := len(p.heap) - 1
	for i := lastChild; i > 0; i /= 2 {
		if p.heap[i/2].Priority <= p.heap[i].Priority {
			break
		}
		p.heap[i], p.heap[i/2] = p.heap[i/2], p.heap[i]
	}
}

func (p *PQ) Dequeue() interface{} {
	lastChild := len(p.heap) - 1
	dequeued := p.heap[0]
	p.heap[0] = p.heap[lastChild]
	p.heap = p.heap[:lastChild]
	minHeapify(p.heap, 0)
	return dequeued.Data
}

func minHeapify(arr []Node, root int) {
	var min int
	size := len(arr)
	left, right := 2*root+1, 2*root+2
	if left < size && arr[left].Priority < arr[root].Priority {
		min = left
	}
	if left < size && arr[right].Priority < arr[root].Priority {
		min = right
	}
	if min != root {
		arr[root], arr[min] = arr[min], arr[root]
		minHeapify(arr, min)
	}
}
