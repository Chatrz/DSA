package heapsort

import maxheap "github.com/sinamna/DSA/Problems/MaxHeapify"

func Sort(arr []int) {
	maxheap.ToMaxHeap(arr)
	size := len(arr)
	for size > 1 {
		size--
		arr[0], arr[size] = arr[size], arr[0]
		maxheap.MaxHeapify(arr[:size], 0)
	}
}
