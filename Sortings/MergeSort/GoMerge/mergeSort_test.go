package GoMerge

import (
	"sort"
	"testing"
)

func TestTestSort(t *testing.T) {
	arr := TestSort()
	if sort.IntsAreSorted(arr)==false{
		t.Error("ARRAY IS NOT SORTED ! ")
	}
}