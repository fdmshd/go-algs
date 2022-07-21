package sorts

func ShellSort(arr []int) {
	n := len(arr)
	for s := n / 2; s > 0; s /= 2 {
		for i := s; i < n; i++ {
			for j := i - s; j >= 0 && arr[j] > arr[j+s]; j -= s {
				arr[j], arr[j+s] = arr[j+s], arr[j]
			}
		}
	}
}
