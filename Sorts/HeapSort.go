package sorts

type Heap struct {
	a        []int
	heapSize int
}

func (h *Heap) Parent(i int) int {
	i = i + 1
	parent := i/2 - 1
	return parent
}

func (h *Heap) Left(i int) int {
	left := 2*i + 1
	return left
}

func (h *Heap) Right(i int) int {
	right := 2*i + 2
	return right
}

func (h *Heap) maxHeapify(i int) {
	l := h.Left(i)
	r := h.Right(i)
	var largest int
	if l < h.heapSize && h.a[l] > h.a[i] {
		largest = l
	} else {
		largest = i
	}
	if r < h.heapSize && h.a[r] > h.a[largest] {
		largest = r
	}
	if largest != i {
		h.a[i], h.a[largest] = h.a[largest], h.a[i]
		h.maxHeapify(largest)
	}
}

func (h *Heap) buildMaxHeap() {
	h.heapSize = len(h.a)
	for i := len(h.a)/2 - 1; i >= 0; i-- {
		h.maxHeapify(i)
	}
}

func Heapsort(arr []int) {
	h := &Heap{
		a: arr,
	}
	h.buildMaxHeap()
	for i := len(arr) - 1; i > 0; i-- {
		h.a[0], h.a[i] = h.a[i], h.a[0]
		h.heapSize -= 1
		h.maxHeapify(0)
	}
}
