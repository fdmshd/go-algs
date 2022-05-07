package sorts

import "sync"

const m = 10

func QuickSort(arr []int, l, r int) {
	if l < r {
		q := partition(arr, l, r)

		QuickSort(arr, l, q)
		QuickSort(arr, q+1, r)
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

func QuickSortPar(arr []int, depth int) {
	if len(arr) <= m {
		InsertionSort(arr)
		return
	}
	l := 0
	r := len(arr) - 1
	i := partition1(arr[l : r+1])
	if depth > 8 {
		QuickSortPar(arr[l:i+1], depth+1)
		QuickSortPar(arr[i+1:r+1], depth+1)
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		QuickSortPar(arr[l:i+1], depth+1)
		wg.Done()
	}()

	go func() {
		QuickSortPar(arr[i+1:r+1], depth+1)
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
