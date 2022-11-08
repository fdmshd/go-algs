package sorts

func CountingSortSimple(arr []int, m int) {
	n := len(arr)
	c := make([]int, m)
	for i := 0; i < n; i++ {
		c[arr[i]]++
	}
	pos := 0
	for i := 0; i < m; i++ {
		for j := 0; j < c[i]; j++ {
			arr[pos] = i
			pos++
		}
	}
}

// broken
func CountingSort(arr []int, m int) []int {
	n := len(arr)
	b := make([]int, n)
	p := make([]int, m)
	for i := 0; i < n; i++ {
		p[arr[i]]++
	}
	for i := 1; i < m; i++ {
		p[i] += p[i-1]
	}
	for i := m - 1; i > 0; i-- {
		p[i] = p[i-1]
	}
	p[0] = 0
	for i := 0; i < n; i++ {
		b[p[arr[i]]] = arr[i]
		p[arr[i]]++
	}
	return b
}
