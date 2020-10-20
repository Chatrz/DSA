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

func TestIsBalanced(t *testing.T) {
	testCases := []struct {
		values []int
		res    bool
	}{
		{
			[]int{5, 3, 4, 7, 8, 6},
			true,
		},
		{
			[]int{1, 2, 3, 4, 5, 6},
			false,
		},
	}
	for _, v := range testCases {
		bst := NewBST()
		bst.Insert(v.values...)
		bst.Print()
		if bst.IsBalanced() != v.res {
			t.Errorf("there is a bug :(")
		}
	}
}
