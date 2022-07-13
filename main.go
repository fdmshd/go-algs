package main

import (
	sorts "algs/Sorts"
	"fmt"
)

func main() {
	arr := []int{5, 6, 7, 1, 9, 8, 2, 3, 4}
	sorts.Heapsort(arr)
	fmt.Println(arr)
}
