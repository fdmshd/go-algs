package sorts

func MergeSort(arr []int) []int {

	lenght := len(arr)
	if lenght >= 2 {
		mid := lenght / 2
		arr = merge(MergeSort(arr[:mid]), MergeSort(arr[mid:]))
	}
	return arr
}
func merge(left []int, right []int) []int {
	//Merge two lists in ascending order.
	lst := make([]int, 0)
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			lst = append(lst, left[0])
			left = left[1:]
		} else {
			lst = append(lst, right[0])
			right = right[1:]
		}
	}
	if len(left) > 0 {
		lst = append(lst, left...)
	}
	if len(right) > 0 {
		lst = append(lst, right...)
	}

	return lst
}
