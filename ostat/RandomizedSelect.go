package ostat

import (
	"math/rand"
	"time"
)

func RandomizedSelect(arr []int, i int) int {
	p := 0
	r := len(arr)
	if len(arr) == 1 {
		return arr[0]
	}
	q := randomizedPartition(arr[p:r])
	k := q - p + 1
	if i == k {
		return arr[q]
	}
	if i < k {
		return RandomizedSelect(arr[p:q], i)
	} else {
		return RandomizedSelect(arr[q+1:r], i-k)
	}

}

func randomizedPartition(arr []int) int {
	r := len(arr) - 1
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(r + 1)
	arr[r], arr[i] = arr[i], arr[r]
	return partition(arr)
}

func partition(a []int) int {
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
