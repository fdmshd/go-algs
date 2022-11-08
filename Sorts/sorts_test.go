package sorts

import (
	"reflect"
	"sort"
	"testing"
	"testing/quick"
)

func TestSort(t *testing.T) {
	fn := func(arr []int) bool {
		cp := make([]int, len(arr))
		copy(cp, arr)
		sort.Ints(cp)
		InsertionSort(arr)
		return reflect.DeepEqual(cp, arr)
	}
	if err := quick.Check(fn, &quick.Config{MaxCount: 2000000}); err != nil {
		t.Error(err)
	}
}
