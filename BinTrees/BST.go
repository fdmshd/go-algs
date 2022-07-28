package bst

import "fmt"

type BSTNode struct {
	Val   int
	P     *BSTNode
	Right *BSTNode
	Left  *BSTNode
}

func NewNode(parent *BSTNode, val int) *BSTNode {
	newNode := &BSTNode{Val: val}
	if parent != nil {
		newNode.P = parent
	}
	return newNode
}

func (node *BSTNode) InorderTraversal() {
	if node == nil {
		return
	}
	node.Left.InorderTraversal()
	fmt.Println(node.Val)
	node.Right.InorderTraversal()
}

func (node *BSTNode) Search(key int) *BSTNode {
	if node == nil || node.Val == key {
		return node
	}

	if key < node.Val {
		return node.Left.Search(key)
	}
	return node.Right.Search(key)
}

func (node *BSTNode) SearchIterative(key int) *BSTNode {
	x := node
	for x != nil && x.Val != key {
		if key < x.Val {
			x = x.Left
		} else {
			x = x.Right
		}
	}
	return x
}

func (node *BSTNode) Min() *BSTNode {
	x := node
	for x.Left != nil {
		x = x.Left
	}
	return x
}

func (node *BSTNode) Max() *BSTNode {
	x := node
	for x.Right != nil {
		x = x.Right
	}
	return x
}

func (node *BSTNode) Successor() *BSTNode {
	if node.Right != nil {
		return node.Right.Min()
	}
	y := node.P
	for y != nil && node == y.Right {
		node = y
		y = y.P
	}
	return y
}

func (node *BSTNode) Predecessor() *BSTNode {
	if node.Left != nil {
		return node.Left.Max()
	}
	y := node.P
	for y != nil && node == y.Left {
		node = y
		y = y.P
	}
	return y
}

func (node *BSTNode) Insert(key int) {
	if key < node.Val {
		if node.Left != nil {
			node.Left.Insert(key)
		} else {
			node.Left = &BSTNode{Val: key}
			node.Left.P = node.Left
		}
	} else {
		if node.Right != nil {
			node.Right.Insert(key)
		} else {
			node.Right = &BSTNode{Val: key}
			node.Right.P = node.Right
		}
	}
}

func (node *BSTNode) Delete(key int) *BSTNode {
	if key < node.Val {
		node.Left = node.Left.Delete(key)
	} else if key > node.Val {
		node.Right = node.Right.Delete(key)
	} else if node.Right != nil && node.Left != nil {
		node.Val = node.Right.Min().Val
		node.Right = node.Right.Delete(node.Val)
	} else {

		if node.Left != nil {
			node = node.Left
		} else if node.Right != nil {
			node = node.Right
		} else {
			node = nil
		}
	}
	return node
}

func (node *BSTNode) RotateLeft() *BSTNode {
	right := node.Right
	right.P = nil
	node.Right = right.Left
	right.Left.P = node
	right.Left = node
	node.P = right
	return right
}

func (node *BSTNode) RotateRight() *BSTNode {
	left := node.Left
	left.P = nil
	node.Left = left.Right
	node.Left.P = node
	left.Right = node
	node.P = left
	return left
}
