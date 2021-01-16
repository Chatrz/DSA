package maxheap

import "testing"

func TestMaxHeapify(t *testing.T) {
	tests := []struct {
		given    []int
		expected []int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{3, 2, 1, 4, 5},
		},
	}
	for _, test := range tests {
		MaxHeapify(test.given, 0)
		for k, v := range test.given {
			if v != test.expected[k] {
				t.Errorf("%v != %v\n", v, test.expected[k])
			}
		}
	}
}
func TestToMaxHeapify(t *testing.T) {
	tests := []struct {
		given    []int
		expected []int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{5, 4, 3, 1, 2},
		},
		{
			[]int{21, 15, 17, 25, 23, 30, 24, 28, 13, 12, 20, 19, 10},
			[]int{30, 28, 24, 25, 23, 19, 21, 15, 13, 12, 20, 17, 10},
		},
	}
	for _, test := range tests {
		ToMaxHeap(test.given)
		for k, v := range test.given {
			if v != test.expected[k] {
				t.Errorf("%v != %v\n", v, test.expected[k])
			}
		}
	}
}
