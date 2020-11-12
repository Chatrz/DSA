package maxheap

import "errors"

func ToMaxHeap(arr []int) {
	size := len(arr)
	for i := (size - 1) / 2; i >= 0; i-- {
		MaxHeapify(arr, i)
	}
}

func HeapExtractMax(heap []int) (int, []int) {
	size := len(heap)
	key := heap[0]
	heap[0], heap[size-1] = heap[size-1], heap[0]
	MaxHeapify(heap[:size-1], 0)
	return key, heap[:size-1]
}

func HeapIncreaseKey(heap []int, key, index int) error {
	if heap[index] > key {
		return errors.New("new key is smaller than current key")
	}
	heap[index] = key
	for i := index; i > 0 && heap[i] > heap[parent(i)]; i = parent(i) {
		heap[parent(i)], heap[i] = heap[i], heap[parent(i)]
	}
	return nil
}

func MaxHeapify(heap []int, root int) {
	size := len(heap)
	left := left(root)
	right := right(root)
	largest := root
	if left < size && heap[left] > heap[root] {
		largest = left
	}
	if right < size && heap[right] > heap[largest] {
		largest = right
	}
	if largest != root {
		heap[root], heap[largest] = heap[largest], heap[root]
		MaxHeapify(heap, largest)
	}
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func parent(i int) int {
	return (i - 1) / 2
}
