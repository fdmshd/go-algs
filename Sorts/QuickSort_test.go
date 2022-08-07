package sorts

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

const n int = 1000000

func BenchmarkQuickSortPar(b *testing.B) {
	var arr1 [n]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		r := rand.Intn(n)
		arr1[i] = r
	}
	b.ResetTimer()
	QuickSortPar(arr1[:], 0)
}
func BenchmarkQuickSort(b *testing.B) {
	var arr2 [n]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		r := rand.Intn(n)
		arr2[i] = r
	}
	b.ResetTimer()
	QuickSort(arr2[:], 0, n-1)
}

func BenchmarkQuickSortS(b *testing.B) {
	var arr2 [n]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		r := rand.Intn(n)
		arr2[i] = r
	}
	b.ResetTimer()
	sort.Ints(arr2[:])
}
