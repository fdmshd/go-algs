package trees

import "math"

type SegmentTree struct {
	tree []int
}

func Left(i int) int {
	return 2*i + 1
}

func Right(i int) int {
	return 2*i + 2
}

func NewSegmentTree(arr []int) SegmentTree {
	n := len(arr)
	newArr := make([]int, int(math.Pow(2, math.Ceil(math.Log2(float64(n))))))
	var j int
	for j = 0; j < n; j++ {
		newArr[j] = arr[j]
	}
	segTree := SegmentTree{
		tree: make([]int, len(newArr)*2-1),
	}
	segTree.buildTree(newArr)
	return segTree
}

func (st *SegmentTree) buildTree(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		st.tree[n+i-1] = arr[i]
	}
	for i := n - 2; i >= 0; i-- {
		st.tree[i] = st.tree[Left(i)] + st.tree[Right(i)]
	}
}

func (st *SegmentTree) Set(i, v int) {
	st.setH(i, v, 0, 0, len(st.tree)/2+1)
}

func (st *SegmentTree) setH(i, v, x, xl, xr int) {
	if xl == xr-1 {
		st.tree[x] = v
		return
	}
	xm := (xl + xr) / 2
	if i < xm {
		st.setH(i, v, Left(x), xl, xm)
	} else {
		st.setH(i, v, Right(x), xm, xr)
	}
	st.tree[x] = st.tree[Left(x)] + st.tree[Right(x)]
}

func (st *SegmentTree) Sum(l, r int) int {
	if l == r {
		return st.tree[len(st.tree)/2+l]
	}
	return st.sumH(l, r+1, 0, 0, len(st.tree)/2+1)
}

func (st *SegmentTree) sumH(l, r, x, xl, xr int) int {
	if xl >= l && xr <= r {
		return st.tree[x]
	}
	if xl >= r || l >= xr {
		return 0
	}
	xm := (xl + xr) / 2
	return st.sumH(l, r, Left(x), xl, xm) + st.sumH(l, r, Right(x), xm, xr)
}
