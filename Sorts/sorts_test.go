package sorts

import (
	"math/rand"
	"sort"
	"testing"
)

const arrSize = 100000

func TestInsertionSort(t *testing.T) {
	arr := generateArray(arrSize)
	InsertionSort(arr)
	if !sort.IntsAreSorted(arr) {
		t.Errorf("array is not sorted")
	}
}

func TestShellSort(t *testing.T) {
	arr := generateArray(arrSize)
	ShellSort(arr)
	if !sort.IntsAreSorted(arr) {
		t.Errorf("array is not sorted")
	}
}

func TestSelectionSort(t *testing.T) {
	arr := generateArray(arrSize)
	SelectionSort(arr)
	if !sort.IntsAreSorted(arr) {
		t.Errorf("array is not sorted")
	}
}

func TestMergeSort(t *testing.T) {
	arr := generateArray(arrSize)
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
	arr := generateArray(arrSize)
	HeapSort(arr)
	if !sort.IntsAreSorted(arr) {
		t.Errorf("array is not sorted")
	}
}

func TestCombSort(t *testing.T) {
	arr := generateArray(arrSize)
	CombSort(arr)
	if !sort.IntsAreSorted(arr) {
		t.Errorf("array is not sorted")
	}
}

func TestCountingSort(t *testing.T) {
	m := 100
	t.Run("CountingSort", func(t *testing.T) {
		arr := generateArrayMax(arrSize, m)
		arr = CountingSort(arr, m)
		if !sort.IntsAreSorted(arr) {
			t.Errorf("array is not sorted")
		}
	})
	t.Run("CountingSort", func(t *testing.T) {
		arr := generateArrayMax(arrSize, m)
		CountingSortSimple(arr, m)
		if !sort.IntsAreSorted(arr) {
			t.Errorf("array is not sorted")
		}
	})
}

func TestQuickSort(t *testing.T) {
	t.Run("QuickSort", func(t *testing.T) {
		arr := generateArray(arrSize)
		QuickSort(arr)
		if !sort.IntsAreSorted(arr) {
			t.Errorf("array is not sorted")
		}
	})
	t.Run("QuickSortPar", func(t *testing.T) {
		arr := generateArray(arrSize)
		QuickSortPar(arr)
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
