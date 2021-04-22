package linked_list

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func Build(s []int) *ListNode {
	var head *ListNode
	for i := len(s) - 1; i >= 0; i-- {
		head = &ListNode{
			s[i],
			head,
		}
	}
	return head
}

func Copy(p *ListNode) *ListNode {
	np := &ListNode{}
	head := np

	for p != nil {
		np.Next = &ListNode{
			Val:  p.Val,
			Next: p.Next,
		}
		p = p.Next
		np = np.Next
	}
	return head.Next
}

func Format(p *ListNode) string {
	s := "["
	for p != nil {
		s += fmt.Sprintf("%d -> ", p.Val)
		p = p.Next
	}
	s += "nil]"
	return s
}
