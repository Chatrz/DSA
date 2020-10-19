package insertionsort

import "testing"

func TestSort(T *testing.T) {
	arr := [...]int{5, 3, 10, 2, 1}
	res := [...]int{1, 2, 3, 5, 10}
	Sort(arr[:])
	if res != arr {
		T.Errorf("%v != %v", arr, res)
	}
}
