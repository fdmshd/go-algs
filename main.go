package main

import (
	sorts "algs/Sorts"
	"fmt"
)

func main() {
	arr := []int{8, 5, 4, 2, 1, 3, 6, 9, 7}
	sorts.CombSort(arr)
	fmt.Println(arr)
}
