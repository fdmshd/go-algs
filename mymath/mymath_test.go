package mymath

import (
	"math"
	"math/rand"
	"testing"
)

func TestMinmax(t *testing.T) {

	wantMax := 101
	wantMin := -1
	t.Run("array size is even", func(t *testing.T) {
		arr := generateArray(1000, 100)
		arr[rand.Intn(500)] = wantMax
		arr[500+rand.Intn(500)] = wantMin

		if gotMin, gotMax := Minmax(arr); gotMin != wantMin || gotMax != wantMax {
			t.Errorf("got: %d %d, want: %d %d", gotMin, gotMax, wantMin, wantMax)
		}
	})

	t.Run("array size is odd", func(t *testing.T) {
		arr := generateArray(999, 100)
		arr[rand.Intn(500)] = wantMax
		arr[500+rand.Intn(499)] = wantMin

		if gotMin, gotMax := Minmax(arr); gotMin != wantMin || gotMax != wantMax {
			t.Errorf("got: %d %d, want: %d %d", gotMin, gotMax, wantMin, wantMax)
		}
	})

}

func generateArray(size, max int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(max)
	}
	return arr
}

func TestMax(t *testing.T) {
	tests := []struct {
		inpA int
		inpB int
		want int
	}{
		{12, 11, 12},
		{-123, 123, 123},
		{1377, 1377, 1377},
		{math.MaxInt, math.MaxInt - 1, math.MaxInt},
		{math.MinInt, math.MinInt + 1, math.MinInt + 1},
		{math.MinInt, math.MaxInt, math.MaxInt},
	}

	for _, test := range tests {
		if got := MaxInt(test.inpA, test.inpB); got != test.want {
			t.Errorf("got: %d, want: %d", got, test.want)
		}
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		inpA int
		inpB int
		want int
	}{
		{12, 11, 11},
		{-123, 123, -123},
		{1377, 1377, 1377},
		{math.MaxInt, math.MaxInt - 1, math.MaxInt - 1},
		{math.MinInt, math.MinInt + 1, math.MinInt},
		{math.MinInt, math.MaxInt, math.MinInt},
	}

	for _, test := range tests {
		if got := MinInt(test.inpA, test.inpB); got != test.want {
			t.Errorf("got: %d, want: %d", got, test.want)
		}
	}
}

func TestABS(t *testing.T) {
	tests := []struct {
		inp  int
		want int
	}{
		{11, 11},
		{-1237, 1237},
		{math.MaxInt, math.MaxInt},
		{math.MinInt + 1, math.MaxInt},
	}

	for _, test := range tests {
		if got := ABSInt(test.inp); got != test.want {
			t.Errorf("got: %d, want: %d", got, test.want)
		}
	}
}
