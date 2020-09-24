package heap

import (
	"github.com/bournex/basic_training/structures/base"
)

// 堆数据结构实现
// 最简单直接的堆实现方式，就是用线型数据结构。
// 通过对n执行2n、2n+1的算数操作，可以很容易找到其子节点的位置。
// 大顶堆的特点是，父节点的值要大于子节点的值。而小顶堆正好相反。
// 由于其使用紧凑的线性空间表达，所以堆天然就是一棵完全二叉树。

type VHeap interface {
	Insert(base.Comparable) error
	ExtraMax() (base.Comparable, error)
	ToString() string
}

func MakeHeap(maxsize int) VHeap {
	h := &heap{maxsize: maxsize + 1}
	h.buf = make([]base.Comparable, h.maxsize)
	return h
}

type heap struct {
	maxsize int // 大小
	len     int
	buf     []base.Comparable
}

func (h *heap) parent(index int) int {
	return index >> 1
}

func (h *heap) left(index int) int {
	return index << 1
}

func (h *heap) right(index int) int {
	return (index << 1) + 1
}

func (h *heap) exchange(i, j int) {
	v := h.buf[i]
	h.buf[i] = h.buf[j]
	h.buf[j] = v
}

// 上浮节点
func (h *heap) up(n int) {

	if n > 1 {
		k := h.parent(n)
		// 只要比父节点大，就继续上浮
		if h.buf[n].Compare(h.buf[k]) > 0 {
			h.exchange(n, k)
			h.up(k)
		}
	}
}

// 下沉节点
func (h *heap) down(n int) {
	l, r := h.left(n), h.right(n)
	max := n
	if (l < h.len+1) && (h.buf[max].Compare(h.buf[l]) < 0) {
		max = l
	}
	if (r < h.len+1) && (h.buf[max].Compare(h.buf[r]) < 0) {
		max = r
	}
	if max != n {
		// 只要有子节点比n节点大，就继续下沉
		h.exchange(n, max)
		h.down(max)
	}
}

// Insert 将新增数据添加到数组的最后，然后对新增数据执行上浮
// 令其上浮到合理位置
func (h *heap) Insert(v base.Comparable) error {
	if h.len < h.maxsize {
		h.buf[h.len+1] = v
		h.up(h.len + 1)
		h.len++
		return nil
	}
	return base.ErrFull
}

// ExtraMax 提取堆中最大值，即根节点
func (h *heap) ExtraMax() (base.Comparable, error) {

	if h.len == 1 {
		// 只有一个节点，返回根节点后，堆被清空
		v := h.buf[1]
		h.buf[1] = nil
		h.len = 0
		return v, nil
	} else if h.len > 1 {
		// 用位节点填充父节点位置，并使其下沉到合适位置
		v := h.buf[1]
		h.buf[1] = h.buf[h.len]
		h.down(1)
		h.len--
		return v, nil
	}
	// 空堆
	return nil, base.ErrEmpty
}

func (h *heap) ToString() string {
	var s string
	for i := 1; i < h.len+1; i++ {
		s += "->"
		s += h.buf[i].ToString()
	}
	return s
}
