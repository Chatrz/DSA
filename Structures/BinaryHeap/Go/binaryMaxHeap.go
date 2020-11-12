////////////////////////////////
//    Author : Armin Goodarzi //
//    GitHub :                //
//      github.com/Armingodiz //
//                            //
////////////////////////////////
package maxHeap

import (
	"errors"
	"fmt"
)

type MaxHeap struct {
	arr         []int
	size        int
	maxCapacity int
}

func NewMaxHeap(capacity int) *MaxHeap {
	heap := &MaxHeap{
		arr:         make([]int, capacity),
		size:        0,
		maxCapacity: capacity,
	}
	return heap
}

func (heap *MaxHeap) maxHeapify(n int, root int) {
	l := left(root)
	r := right(root)
	largest := root
	if l < n && heap.arr[l] > heap.arr[root] {
		largest = l
	}
	if r < n && heap.arr[r] > heap.arr[largest] {
		largest = r
	}
	if largest != root {
		heap.exchange(root, largest)
		heap.maxHeapify(n, largest)
	}
}

func buildHeap(arr []int, capacity int) *MaxHeap {
	heap := &MaxHeap{arr: arr, size: len(arr), maxCapacity: capacity}
	n := heap.size
	// Build heap (rearrange array)
	for i := n/2 - 1; i >= 0; i-- {
		heap.maxHeapify(n, i)
	}
	return heap
}

func heapSort(arr []int, capacity int) []int {
	heap := buildHeap(arr, capacity)
	n := heap.size
	// One by one extract an element from heap
	for i := n - 1; i > 0; i-- {
		// Move current root to end
		heap.exchange(0, i)
		// call max heapify on the reduced heap
		heap.maxHeapify(i, 0)
	}
	return heap.arr
}

func parent(index int) int {
	return (index - 1) / 2
}

func left(index int) int {
	return 2*index + 1
}

func right(index int) int {
	return 2*index + 2
}

func (heap *MaxHeap) GetMax() int {
	return heap.arr[0]
}

func (maxHeap *MaxHeap) Insert(key int) error {
	if maxHeap.size >= maxHeap.maxCapacity {
		return errors.New("heap is full !")
	} else {
		// inserting key in heap
		index := maxHeap.size
		maxHeap.arr[index] = key
		maxHeap.size += 1
		// maintaining the heap after adding key
		for index != 0 && maxHeap.arr[index] > maxHeap.arr[parent(index)] {
			maxHeap.exchange(index, parent(index))
			index = parent(index)
		}
	}
	return nil
}

func (heap *MaxHeap) IncreaseKey(index, newKey int) {
	heap.arr[index] = newKey
	for index != 0 && heap.arr[index] > heap.arr[parent(index)] {
		heap.exchange(index, parent(index))
		index = parent(index)
	}
}

func (heap *MaxHeap) ExtractMax() int {
	tmp := heap.arr[0]
	heap.arr[0] = heap.arr[heap.size-1]
	heap.size -= 1
	heap.maxHeapify(heap.size, 0)
	return tmp
}

func (heap *MaxHeap) DeleteKey(index int) {
	heap.IncreaseKey(index, heap.GetMax())
	heap.ExtractMax()
}

func (maxHeap *MaxHeap) exchange(index1, index2 int) {
	tmp := maxHeap.arr[index1]
	maxHeap.arr[index1] = maxHeap.arr[index2]
	maxHeap.arr[index2] = tmp
}

func (maxHeap *MaxHeap) printHeap() {
	fmt.Println("HEAP ELEMENTS : ")
	for i := 0; i < maxHeap.size; i++ {
		fmt.Print(maxHeap.arr[i])
		fmt.Print(" ")
	}
	fmt.Println()
}

/*
  func main() {
	heap := NewMaxHeap(10)
	heap.Insert(1)
	heap.Insert(3)
	heap.Insert(4)
	heap.Insert(10)
	heap.Insert(2)
	heap.Insert(19)
	heap.Insert(0)
	heap.Insert(87)
	heap.printHeap()
  fmt.Println("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&7")
  heap.DeleteKey(6)
  heap.printHeap()
  heap.DeleteKey(0)
  heap.printHeap()
  heap.DeleteKey(1)
  heap.printHeap()
  fmt.Println("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&7")
	array := []int{12, 11, 13, 5, 6, 7}
	array = heapSort(array, 10)
	fmt.Println(array)
	newHeap := buildHeap(array, 10)
	newHeap.printHeap()
}
*/
