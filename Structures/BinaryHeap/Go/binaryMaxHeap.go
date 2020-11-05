package main

import (
	"errors"
	"fmt"
)

type MaxHeap struct {
	arr         []int
	size        int
	maxCapacity int
}

type arrHolder struct {
	arr []int
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

func buildHeap(arr []int,capacity int) *MaxHeap{
  heap := &MaxHeap{arr: arr, size: len(arr), maxCapacity: capacity}
  n := heap.size
  // Build heap (rearrange array)
  for i := n/2 - 1; i >= 0; i-- {
    heap.maxHeapify(n, i)
  }
  return heap
}

func heapSort(arr []int, capacity int) []int {
	heap := buildHeap(arr,capacity)
  n:=heap.size
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
	fmt.Println("maxheapify array : ")
	array := []int{12, 11, 13, 5, 6, 7}
	array = heapSort(array , 10)
  fmt.Println(array)
  newHeap :=buildHeap(array,10)
  newHeap.printHeap()
}
