package main

import (
	bst "algs/BinTrees"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	root := &bst.AVLNode{Val: 5}
	var input int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		input = rand.Intn(20)
		root = root.Insert(input)
	}

	fmt.Println(root)

}
