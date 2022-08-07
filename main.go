package main

import (
	"algs/trees"
	"fmt"
)

func main() {
	arr := []int{0, 9, 5, 7, 3}
	tree := trees.NewSegmentTree(arr)
	fmt.Println(tree)
	fmt.Println(tree.Sum(2, 4))
}
