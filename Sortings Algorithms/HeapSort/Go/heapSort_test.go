package heapsort

import "testing"

func TestSort(t *testing.T) {
	tests := []struct {
		given    []int
		expected []int
	}{
		{
			[]int{5, 4, 3, 2, 1},
			[]int{1, 2, 3, 4, 5},
		},
	}
	for _, test := range tests {
		Sort(test.given)
		for k, v := range test.given {
			if v != test.expected[k] {
				t.Errorf("%v != %v\n", v, test.expected[k])
			}
		}
	}
}
