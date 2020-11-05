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

func NewMaxHeap(capacity int) *MaxHeap {
	heap := &MaxHeap{
		arr:         make([]int, capacity),
		size:        0,
		maxCapacity: capacity,
	}
	return heap
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
}
