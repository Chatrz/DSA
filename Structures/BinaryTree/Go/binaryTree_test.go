package binarytree

import "testing"

func TestDepth(t *testing.T) {
	testCases := []struct {
		values []int
		res    int
	}{
		{
			[]int{5, 3, 4, 7, 8, 6},
			3,
		},
	}
	for _, v := range testCases {
		bt := NewBT()
		bt.Insert(v.values...)
		bt.Print()
		if bt.Depth() != v.res {
			t.Errorf("%v != %v", bt.Depth(), v.res)
		}
	}
}
