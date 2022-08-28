package search

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateArray(n int) []int {
	m := make(map[int]bool)
	arr := make([]int, n)
	i := 0
	for i < n {
		newElem := rand.Intn(100)
		if !m[newElem] {
			arr[i] = newElem
			m[newElem] = true
			i++
		}
	}
	sort.Ints(arr)
	return arr
}

func TestBinSearchGen(t *testing.T) {
	arr := generateArray(20)
	keyI := rand.Intn(20)
	t.Logf("key %d in array %v", arr[keyI], arr)
	if keyI != BinSearchRecursion(arr, arr[keyI], 0, len(arr)-1) {
		t.Errorf("Элемент %d расположен в позиции %d", arr[keyI], keyI)
	}
	if keyI != BinSearchLoop(arr, arr[keyI]) {
		t.Errorf("Элемент %d расположен в позиции %d", arr[keyI], keyI)
	}
}

func TestBinSearchExtreme(t *testing.T) {
	arr := []int{2, 5, 6, 7, 8, 9, 11, 21}
	key := 2
	res := BinSearchRecursion(arr, key, 0, len(arr)-1)
	exp := 0
	if res != exp {
		t.Errorf("Для BinSearchRecursion с ключем %d в массиве %v: act=%d, exp=%d",
			key, arr, res, exp)
	}
	res = BinSearchLoop(arr, key)
	if res != exp {
		t.Errorf("Для BinSearchLoop с ключем %d в массиве %v: act=%d, exp=%d",
			key, arr, res, exp)
	}

	key = 21
	res = BinSearchRecursion(arr, key, 0, len(arr)-1)
	exp = 7
	if res != exp {
		t.Errorf("Для BinSearchRecursion с ключем %d в массиве %v: act=%d, exp=%d",
			key, arr, res, exp)
	}
	res = BinSearchLoop(arr, key)
	if res != exp {
		t.Errorf("Для BinSearchLoop с ключем %d в массиве %v: act=%d, exp=%d",
			key, arr, res, exp)
	}
}

func TestBinSearchNeg(t *testing.T) {
	arr := []int{2, 5, 6, 7, 8, 9, 11, 21}
	key := 3
	res := BinSearchRecursion(arr, key, 0, len(arr)-1)
	exp := -1
	if res != exp {
		t.Errorf("Для BinSearchRecursion с ключем %d в массиве %v: act=%d, exp=%d",
			key, arr, res, exp)
	}
	res = BinSearchLoop(arr, key)
	if res != exp {
		t.Errorf("Для BinSearchLoop с ключем %d в массиве %v: act=%d, exp=%d",
			key, arr, res, exp)
	}
}

func TestInterpolationSearch(t *testing.T) {
	arr := []int{2, 5, 6, 7, 8, 9, 11, 21}
	key := 21
	exp := 7
	res := InterSearch(arr, key)
	if res != exp {
		t.Errorf("arr=%v, key=%d, res= %d, exp=%d", arr, key, res, exp)
	}

	key = 7
	exp = 3
	res = InterSearch(arr, key)
	if res != exp {
		t.Errorf("arr=%v, key=%d, res= %d, exp=%d", arr, key, res, exp)
	}
}
