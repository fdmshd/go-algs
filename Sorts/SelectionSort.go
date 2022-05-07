package sorts

func SelectionSort(arr []int) {
	var min int
	for i := 0; i < len(arr)-1; i++ {
		min = findMin(arr, i)
		arr[i], arr[min] = arr[min], arr[i]
	}
}
func findMin(arr []int, l int) int {
	var min int
	min = l
	for i := l + 1; i < len(arr); i++ {
		if arr[i] < arr[min] {
			min = i
		}
	}
	return min
}
