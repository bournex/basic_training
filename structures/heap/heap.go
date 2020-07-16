package heap

import (
	"github.com/bournex/basic_training/structures/base"
)

// 堆数据结构实现

type VHeap interface {
	Insert(base.Comparable) error
	ExtraMax() (base.Comparable, error)
	ToString() string
}

func MakeHeap(maxsize int) VHeap {
	h := &Heap{maxsize: maxsize + 1}
	h.heap = make([]base.Comparable, h.maxsize)
	return h
}

type Heap struct {
	maxsize int // 大小
	len     int
	heap    []base.Comparable
}

func (h *Heap) parent(index int) int {
	return index >> 1
}

func (h *Heap) left(index int) int {
	return index << 1
}

func (h *Heap) right(index int) int {
	return (index << 1) + 1
}

func (h *Heap) exchange(i, j int) {
	v := h.heap[i]
	h.heap[i] = h.heap[j]
	h.heap[j] = v
}

func (h *Heap) up(n int) {

	if n > 1 {
		k := h.parent(n)
		if h.heap[n].Compare(h.heap[k]) > 0 {
			h.exchange(n, k)
			h.up(k)
		}
	}
}

func (h *Heap) down(n int) {
	l, r := h.left(n), h.right(n)
	max := n
	if (l < h.len+1) && (h.heap[max].Compare(h.heap[l]) < 0) {
		max = l
	}
	if (r < h.len+1) && (h.heap[max].Compare(h.heap[r]) < 0) {
		max = r
	}
	if max != n {
		h.exchange(n, max)
		h.down(max)
	}
}

func (h *Heap) Insert(v base.Comparable) error {
	if h.len < h.maxsize {
		h.heap[h.len+1] = v
		h.up(h.len + 1)
		h.len++
		return nil
	}
	return base.ErrFull
}

func (h *Heap) ExtraMax() (base.Comparable, error) {

	if h.len == 1 {
		v := h.heap[1]
		h.heap[1] = nil
		h.len = 0
		return v, nil
	} else if h.len > 1 {
		v := h.heap[1]
		h.heap[1] = h.heap[h.len]
		h.down(1)
		h.len--
		return v, nil
	}
	return nil, base.ErrEmpty
}

func (h *Heap) ToString() string {
	var s string
	for i := 1; i < h.len+1; i++ {
		s += "->"
		s += h.heap[i].ToString()
	}
	return s
}
