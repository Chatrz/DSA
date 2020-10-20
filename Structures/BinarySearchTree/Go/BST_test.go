package binarysearchtree

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
		bst := NewBST()
		bst.Insert(v.values...)
		bst.Print()
		if bst.Depth() != v.res {
			t.Errorf("%v != %v", bst.Depth(), v.res)
		}
	}
}
