package MaxHeapify

func ToMaxHeap(arr []int) {
	size := len(arr)
	for i := (size - 1) / 2; i >= 0; i-- {
		MaxHeapify(arr, i)
	}
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
