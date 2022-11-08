package bst

import "algs/mymath"

type AVLNode struct {
	Val    int
	Height int
	Right  *AVLNode
	Left   *AVLNode
}

func (node *AVLNode) BalanceFactor() int {
	rightH := 0
	leftH := 0
	if node.Right != nil {
		rightH = node.Right.Height
	}
	if node.Left != nil {
		leftH = node.Left.Height
	}
	return rightH - leftH
}

func (node *AVLNode) FixHeight() {
	rightH := 0
	leftH := 0
	if node.Right != nil {
		rightH = node.Right.Height
	}
	if node.Left != nil {
		leftH = node.Left.Height
	}
	node.Height = mymath.MaxInt(rightH, leftH) + 1
}

func (node *AVLNode) RotateLeft() *AVLNode {
	newNode := node.Right
	node.Right = newNode.Left
	newNode.Left = node
	node.FixHeight()
	newNode.FixHeight()
	return newNode
}

func (node *AVLNode) RotateRight() *AVLNode {
	newNode := node.Left
	node.Left = newNode.Right
	newNode.Left = node
	node.FixHeight()
	newNode.FixHeight()
	return newNode
}

func (node *AVLNode) Balance() *AVLNode {
	node.FixHeight()
	if node.BalanceFactor() == 2 {
		if node.Right.BalanceFactor() < 0 {
			node.Right = node.Right.RotateRight()
		}
		return node.RotateLeft()
	}

	if node.BalanceFactor() == -2 {
		if node.Left.BalanceFactor() > 0 {
			node.Left = node.Left.RotateLeft()
		}
		return node.RotateRight()
	}
	return node
}

func (node *AVLNode) Insert(key int) *AVLNode {
	if key < node.Val {
		if node.Left != nil {
			node.Left.Insert(key)
		} else {
			node.Left = &AVLNode{Val: key}
		}
	} else {
		if node.Right != nil {
			node.Right.Insert(key)
		} else {
			node.Right = &AVLNode{Val: key}

		}
	}
	return node.Balance()
}

// не лучшая реализация
func (node *AVLNode) Delete(key int) *AVLNode {
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
	return node.Balance()
}

func (node *AVLNode) Min() *AVLNode {
	x := node
	for x.Left != nil {
		x = x.Left
	}
	return x
}
