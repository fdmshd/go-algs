package main

import (
	substrings "algs/Substrings"
	"fmt"
)

func main() {
	s := []rune("привет мир")
	t := []rune("мир")
	fmt.Println(substrings.BoyerMoore(s, t))
}
