package sorts

func InsertionSort(arr []int) {

	for i := 1; i < len(arr); i++ {
		key := arr[i]
		var j int
		for j = i - 1; (j >= 0) && (key < arr[j]); j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = key
	}
}
