package main

import (
	sorts "algs/Sorts"
	"fmt"
)

func main() {
	arr := []int{329, 457, 657, 839, 436, 720, 355}
	sorts.LCD(arr)
	fmt.Println(arr)
}
