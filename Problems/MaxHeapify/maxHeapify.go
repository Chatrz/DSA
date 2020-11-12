package maxheap

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
