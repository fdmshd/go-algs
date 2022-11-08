package sorts

import (
	"math/rand"
	"sort"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	arr := generateArray(100000)
	InsertionSort(arr)
	if !sort.IntsAreSorted(arr) {
		t.Errorf("array is not sorted")
	}
}

func TestShellSort(t *testing.T) {
	arr := generateArray(100000)
	ShellSort(arr)
	if !sort.IntsAreSorted(arr) {
		t.Errorf("array is not sorted")
	}
}

func TestSelectionSort(t *testing.T) {
	arr := generateArray(100000)
	SelectionSort(arr)
	if !sort.IntsAreSorted(arr) {
		t.Errorf("array is not sorted")
	}
}

func TestMergeSort(t *testing.T) {
	arr := generateArray(100000)
	arr = MergeSort(arr)
	if !sort.IntsAreSorted(arr) {
		t.Errorf("array is not sorted")
	}
}

func TestLSD(t *testing.T) {
	arr := generateArray(100)
	LSD(arr)
	if !sort.IntsAreSorted(arr) {
		t.Errorf("array is not sorted")
	}
}

func TestHeapSort(t *testing.T) {
	arr := generateArray(100000)
	HeapSort(arr)
	if !sort.IntsAreSorted(arr) {
		t.Errorf("array is not sorted")
	}
}

func TestCombSort(t *testing.T) {
	arr := generateArray(100000)
	CombSort(arr)
	if !sort.IntsAreSorted(arr) {
		t.Errorf("array is not sorted")
	}
}

func TestCountingSort(t *testing.T) {
	m := 100
	t.Run("CountingSort", func(t *testing.T) {
		arr := generateArrayMax(1000, m)
		arr = CountingSort(arr, m)
		if !sort.IntsAreSorted(arr) {
			t.Errorf("array is not sorted")
		}
	})
	t.Run("CountingSort", func(t *testing.T) {
		arr := generateArrayMax(100000, m)
		CountingSortSimple(arr, m)
		if !sort.IntsAreSorted(arr) {
			t.Errorf("array is not sorted")
		}
	})
}

func generateArrayMax(size, max int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(size)
	}
	return arr
}

func generateArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Int()
	}
	return arr
}
