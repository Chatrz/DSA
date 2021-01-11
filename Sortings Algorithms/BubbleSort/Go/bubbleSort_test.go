package bubblesort

import (
	"fmt"
	"testing"
	"time"
)

func TestSort(T *testing.T) {
	exTime := time.Now().UnixNano()
	arr := [...]int{5, 3, 10, 2, 1}
	res := [...]int{1, 2, 3, 5, 10}
	Sort(arr[:])
	if res != arr {
		T.Errorf("%v != %v", arr, res)
	}
	exTime = time.Now().UnixNano() - exTime
	fmt.Printf("expired time: %v ms", exTime)
}

func TestOptimizedSort(T *testing.T) {
	exTime := time.Now().UnixNano()
	arr := [...]int{5, 3, 10, 2, 1}
	res := [...]int{1, 2, 3, 5, 10}
	OptimizedSort(arr[:])
	if res != arr {
		T.Errorf("%v != %v", arr, res)
	}
	exTime = time.Now().UnixNano() - exTime
	fmt.Printf("expired time with optimization: %v ms", exTime)
}
