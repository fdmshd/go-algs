package bs

func BinSearchRecursion(arr []int, key int, l int, r int) int {
	if key < arr[l] || key > arr[r] {
		return -1
	}
	if r < l {
		return -1
	}
	answer := -1
	middle := (l + r) / 2
	if arr[middle] == key {
		return middle
	}
	if arr[middle] < key {
		answer = BinSearchRecursion(arr, key, middle+1, r)
	} else {
		answer = BinSearchRecursion(arr, key, l, middle-1)
	}
	return answer
}
func BinSearchLoop(arr []int, key int) int {
	l := 0
	r := len(arr) - 1
	if key < arr[l] || key > arr[r] {
		return -1
	}
	var middle int
	for l <= r {
		middle = (l + r) / 2
		if arr[middle] == key {
			return middle
		}
		if key > arr[middle] {
			l = middle + 1
		} else {
			r = middle - 1
		}
	}
	return -1
}
