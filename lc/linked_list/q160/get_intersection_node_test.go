package q160

import (
	"fmt"
	"testing"
)

func genList(array []int) *ListNode {
	head := new(ListNode)
	p := head
	for _, v := range array {
		p.Next = &ListNode{v, nil}
		p = p.Next
	}
	return head.Next
}

func genIntersectionList(array1, array2, array3 []int) (*ListNode, *ListNode, *ListNode) {

	l1 := genList(array1)
	l2 := genList(array2)
	l3 := genList(array3)

	p := l1
	for p != nil {
		if p.Next == nil {
			p.Next = l3
			break
		}
		p = p.Next
	}
	p = l2
	for p != nil {
		if p.Next == nil {
			p.Next = l3
			break
		}
		p = p.Next
	}
	return l1, l2, l3
}
func TestGetIntersectionNode(t *testing.T) {
	type Case struct {
		InputA *ListNode
		InputB *ListNode
		Expect *ListNode
	}

	genCase := func(array1, array2, array3 []int) Case {
		var c Case
		c.InputA, c.InputB, c.Expect = genIntersectionList(array1, array2, array3)
		return c
	}

	examples := []Case{
		genCase([]int{1, 2}, []int{1, 3, 4}, []int{5, 6}),
		genCase([]int{1, 2}, []int{1, 3}, nil),
	}

	for _, v := range examples {
		res := getIntersectionNode(v.InputA, v.InputB)

		if res != v.Expect {
			t.Errorf("input a %s, input b %s, expect %s, got %s", formatList(v.InputA), formatList(v.InputB), formatList(v.Expect), formatList(res))
		}
	}
}

func formatList(p *ListNode) string {
	s := "["
	for p != nil {
		s += fmt.Sprintf("%d -> ", p.Val)
		p = p.Next
	}
	s += "nil]"
	return s
}
