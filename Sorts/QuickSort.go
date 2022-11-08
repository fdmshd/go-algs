package sorts

import (
	"runtime"
	"sync"
)

const m = 10

var maxDepth = runtime.NumCPU() * 2

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, l, r int) {
	if len(arr) <= m {
		InsertionSort(arr)
		return
	}
	if l < r {
		q := partition(arr, l, r)

		quickSort(arr, l, q)
		quickSort(arr, q+1, r)
	}
}
func partition(a []int, l, r int) int {
	x := a[(l+r)/2]
	for l <= r {
		for a[l] < x {
			l++
		}
		for a[r] > x {
			r--
		}
		if l >= r {
			break
		}
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
	return r
}

func QuickSortPar(arr []int) {
	quickSortPar(arr, 0)
}

func quickSortPar(arr []int, depth int) {
	if len(arr) <= m {
		InsertionSort(arr)
		return
	}
	l := 0
	r := len(arr) - 1
	i := partition1(arr[l : r+1])
	if depth > maxDepth {
		quickSortPar(arr[l:i+1], depth+1)
		quickSortPar(arr[i+1:r+1], depth+1)
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		quickSortPar(arr[l:i+1], depth+1)
		wg.Done()
	}()

	go func() {
		quickSortPar(arr[i+1:r+1], depth+1)
		wg.Done()
	}()
	wg.Wait()
}

func partition1(a []int) int {
	l := 0
	r := len(a) - 1
	x := a[(l+r)/2]
	for l <= r {
		for a[l] < x {
			l++
		}
		for a[r] > x {
			r--
		}
		if l >= r {
			break
		}
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
	return r
}
