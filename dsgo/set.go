package main

type Heap[T any] struct {
	data []T
	comp func(a, b T) bool
}

func NewHeap[T any](comp func(a, b T) bool) *Heap[T] {
	return &Heap[T]{
		data: make([]T, 0),
		comp: comp,
	}
}

func (h *Heap[T]) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *Heap[T]) Peek() (T, bool) {
	if h.Size() == 0 {
		var val T
		return val, false
	}
	return h.data[0], true
}

func (h *Heap[T]) Size() int {
	return len(h.data)
}

func leftChildIndex(i int) int {
	return 2*i + 1
}
func rightChildIndex(i int) int {
	return 2*i + 2
}

func parentIndex(i int) int {
	return (i - 1) / 2
}

func (h *Heap[T]) heapifyDown(i int) {
	l, r := leftChildIndex(i), rightChildIndex(i)
	largest := i
	if l < len(h.data) && h.comp(h.data[i], h.data[l]) {
		largest = l
	}
	if r < len(h.data) && h.comp(h.data[largest], h.data[r]) {
		largest = r
	}
	if largest != i {
		h.swap(i, largest)
		h.heapifyDown(largest)
	}
}
func (h *Heap[T]) heapifyUp(i int) {
	for h.comp(h.data[i], h.data[parentIndex(i)]) {
		h.swap(i, parentIndex(i))
		i = parentIndex(i)
	}
}
func (h *Heap[T]) Pop() (T, bool) {
	var val T
	if h.Size() == 0 {
		return val, false
	}
	val = h.data[0]
	h.swap(0, h.Size()-1)
	h.data = h.data[:h.Size()-1]
	h.heapifyDown(0)
	return val, true

}
func (h *Heap[T]) Push(v T) {
	h.data = append(h.data, v)
	h.heapifyUp(h.Size() - 1)
}
func testHeap() {
	h := NewHeap[int](func(a, b int) bool { return a < b })
	h.data = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	h.heapifyDown(0)
}
