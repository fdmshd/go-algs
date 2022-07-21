package sorts

func CombSort(arr []int) {
	n := len(arr)
	factor := 1.2473309
	step := n - 1
	for step >= 1 {
		for i := 0; i+step < n; i++ {
			if arr[i] > arr[i+step] {
				arr[i], arr[i+step] = arr[i+step], arr[i]
			}
		}
		step = int(float64(step) / factor)
	}
}
