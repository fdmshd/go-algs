package search

// Interpolation search
// Некорректно работает =( см. тесты
func InterSearch(arr []int, key int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		m := (key - arr[left]) * (right - left) / (arr[right] - arr[left])
		if arr[m] == key {
			return m
		}
		if arr[m] < key {
			left = m
		} else {
			right = m
		}
	}
	return -1
}
