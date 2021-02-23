package sorting

import (
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	sorted := bubbleSort(arr)
	if sort.IntsAreSorted(sorted) == false {
		t.Errorf("arr is not sorted")
	}
}
