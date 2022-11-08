package mymath

func Minmax(arr []int) (min int, max int) {
	n := len(arr)
	i := 0
	min = arr[i]
	if n%2 != 0 {
		max = arr[i]
		i++
	} else {
		max = arr[i+1]
		i += 2
	}
	for ; i < n-1; i++ {
		var pairMin, pairMax int
		if arr[i] < arr[i+1] {
			pairMin = arr[i]
			pairMax = arr[i+1]
		} else {
			pairMin = arr[i+1]
			pairMax = arr[i]
		}

		if pairMax > max {
			max = pairMax
		}
		if pairMin < min {
			min = pairMin
		}
	}
	return min, max
}

func ABSInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func MaxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func MinInt(a, b int) int {
	if b < a {
		return b
	}
	return a
}
