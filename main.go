package main

import (
	bst "algs/BinTrees"
	"fmt"
)

func main() {
	root := &bst.BSTNode{Val: 1}
	left := &bst.BSTNode{Val: 0, P: root}
	right := &bst.BSTNode{Val: 2, P: root}
	root.Left = left
	root.Right = right
	root.Insert(8)
	root = root.Delete(1)
	fmt.Println(root)

}
