package rbst

import (
	"math/rand"
	"time"
)

type RTreeNode struct {
	Key   int
	Size  int
	Left  *RTreeNode
	Right *RTreeNode
}

func CreateNode(k int) *RTreeNode {
	return &RTreeNode{
		Key:  k,
		Size: 1,
	}
}

func find(node *RTreeNode, key int) *RTreeNode {
	if node == nil {
		return nil
	}
	if key == node.Key {
		return node
	}
	if key < node.Key {
		return find(node.Left, key)
	} else {
		return find(node.Right, key)
	}
}

func getSize(node *RTreeNode) int {
	if node == nil {
		return 0
	}
	return node.Size
}

func fixSize(node *RTreeNode) {
	node.Size = getSize(node.Left) + getSize(node.Right) + 1
}

func rotateRight(p *RTreeNode) *RTreeNode {
	q := p.Left
	if q == nil {
		return p
	}
	p.Left = q.Right
	q.Right = p
	q.Size = p.Size
	fixSize(p)
	return q
}

func rotateLeft(q *RTreeNode) *RTreeNode {
	p := q.Right
	q.Right = p.Left
	p.Left = q
	p.Size = q.Size
	fixSize(q)
	return p
}

func InsertRoot(node *RTreeNode, key int) *RTreeNode {

	if node == nil {
		return &RTreeNode{Key: key}
	}
	if key < node.Key {
		node.Left = InsertRoot(node.Left, key)
		return rotateRight(node)
	} else {
		node.Right = InsertRoot(node.Right, key)
		return rotateLeft(node)
	}

}

func Insert(node *RTreeNode, key int) *RTreeNode {
	if node == nil {
		return &RTreeNode{Key: key}
	}
	rand.Seed(time.Now().UnixNano())

	if rand.Int()%(node.Size+1) == 0 {
		return InsertRoot(node, key)
	}
	if node.Key > key {
		node.Left = Insert(node.Left, key)
	} else {
		node.Right = Insert(node.Right, key)
	}
	fixSize(node)
	return node
}

func join(p, q *RTreeNode) *RTreeNode {
	if p == nil {
		return q
	}
	if q == nil {
		return p
	}
	rand.Seed(time.Now().UnixNano())
	if rand.Int()%(p.Size+q.Size) < p.Size {
		p.Right = join(p.Right, q)
		fixSize(p)
		return p
	} else {
		q.Left = join(p, q.Left)
		fixSize(q)
		return q
	}
}

func remove(p *RTreeNode, key int) *RTreeNode {
	node := find(p, key)
	q := join(node.Left, node.Right)
	return q
}
