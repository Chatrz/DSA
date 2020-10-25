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
		//bst.Print()
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
		//bst.Print()
		if bst.IsBalanced() != v.res {
			t.Errorf("there is a bug :(")
		}
	}
}

func TestSeccessor(t *testing.T) {
	tests := []struct {
		values []int
		res    map[int]int
	}{
		{
			[]int{20, 8, 22, 4, 12, 10, 14},
			map[int]int{8: 10, 10: 12, 14: 20},
		},
	}
	for _, test := range tests {
		bst := NewBST()
		bst.Insert(test.values...)
		for k, v := range test.res {
			node := bst.Search(k)
			if v != node.Seccessor().Value {
				t.Errorf("%v != %v", v, node.Seccessor().Value)
			}
		}
	}
}

func TestPredecessor(t *testing.T) {
	tests := []struct {
		values []int
		res    map[int]int
	}{
		{
			[]int{20, 8, 22, 4, 12, 10, 14},
			map[int]int{8: 4, 10: 8, 14: 12},
		},
	}
	for _, test := range tests {
		bst := NewBST()
		bst.Insert(test.values...)
		for k, v := range test.res {
			node := bst.Search(k)
			if v != node.Predecessor().Value {
				t.Errorf("%v != %v", v, node.Predecessor().Value)
			}
		}
	}
}
