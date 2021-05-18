package h23

import (
	ll "github.com/bournex/basic_training/lc/linked_list"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 思路
//	刚开始做困难题目，没有思路，按照答案实现的，就当是复习优先队列了
func mergeKLists(lists []*ll.ListNode) *ll.ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	} else if n == 1 {
		return lists[0]
	}

	p := createPriority(n)
	for i := 0; i < n; i++ {
		p.add(lists[i])
	}

	head := &ll.ListNode{Next: p.getMin()}
	prev := head
	curr := prev.Next
	for curr != nil {
		if curr.Next != nil {
			p.add(curr.Next)
			curr.Next = nil
		}
		prev.Next = curr
		prev = curr
		curr = p.getMin()
	}

	return head.Next
}

// 最小堆
type priorityQueue struct {
	Nodes []*ll.ListNode
	Len   int // 当前总长度
}

func createPriority(n int) *priorityQueue {
	pq := new(priorityQueue)
	pq.Nodes = make([]*ll.ListNode, n+1)
	pq.Len = 0
	return pq
}

func (p *priorityQueue) getMin() *ll.ListNode {
	if p.Len > 0 {
		p.Nodes[1], p.Nodes[p.Len] = p.Nodes[p.Len], p.Nodes[1]
		p.Len--
		p.down(1)
		return p.Nodes[p.Len+1]
	}
	return nil
}

func (p *priorityQueue) add(node *ll.ListNode) {
	if node != nil {
		p.Nodes[p.Len+1] = node
		p.Len++
		p.up(p.Len)
	}
}

func (p *priorityQueue) up(idx int) {
	parent := idx >> 1
	for parent >= 1 {
		if p.Nodes[parent].Val > p.Nodes[idx].Val {
			p.Nodes[parent], p.Nodes[idx] = p.Nodes[idx], p.Nodes[parent]
		}
		idx = parent
		parent = idx >> 1
	}
}

func (p *priorityQueue) down(idx int) {
	min := idx
	for idx <= p.Len {
		if idx<<1 <= p.Len && p.Nodes[min].Val > p.Nodes[idx<<1].Val {
			min = idx << 1
		}
		if idx<<1+1 <= p.Len && p.Nodes[min].Val > p.Nodes[idx<<1+1].Val {
			min = idx<<1 + 1
		}
		if min == idx {
			break
		}
		p.Nodes[min], p.Nodes[idx] = p.Nodes[idx], p.Nodes[min]
		idx = min
	}
}
