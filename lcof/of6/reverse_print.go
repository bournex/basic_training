package of6

import "container/list"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// slice实现
func reversePrint(head *ListNode) []int {
	n := 0
	curr := head
	for curr != nil {
		n++
		curr = curr.Next
	}

	s := make([]int, n)
	curr = head
	for curr != nil {
		s[n-1] = curr.Val
		n--
		curr = curr.Next
	}
	return s
}

// list实现
func reversePrint1(head *ListNode) []int {
	l := list.New()
	for head != nil {
		l.PushBack(head.Val)
		head = head.Next
	}
	s := make([]int, l.Len())
	for i := 0; i < len(s); i++ {
		e := l.Back()
		s[i] = e.Value.(int)
		l.Remove(e)
	}
	return s[0:]
}
