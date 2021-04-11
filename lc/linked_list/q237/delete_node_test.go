package q237

import (
	"fmt"
	"testing"
)

type Case struct {
	Head   *ListNode
	Input  *ListNode
	Expect *ListNode
}

func genCase(array []int, delIndex int) Case {
	var toDel *ListNode
	dummyInput := &ListNode{0, nil}
	dummyExpect := &ListNode{0, nil}
	input, expect := dummyInput, dummyExpect
	for i, v := range array {
		input.Next = &ListNode{v, nil}
		input = input.Next
		if i != delIndex {
			expect.Next = &ListNode{v, nil}
			expect = expect.Next
		} else {
			toDel = input
		}
	}
	return Case{
		Head:   dummyInput.Next,
		Input:  toDel,
		Expect: dummyExpect.Next,
	}
}

func TestDeleteNode(t *testing.T) {
	examples := []Case{
		genCase([]int{1, 2, 3, 4}, 2), // expect [1,2,4]，delete midian node
		genCase([]int{1, 2, 3}, 0),    // expect[1,2]，delete head
		// genCase([]int{1, 2, 3}, 2),    // expect[1,2]，delete tail
	}

	for _, v := range examples {
		head := copyList(v.Head)
		val := v.Input.Val
		deleteNode(v.Input)

		res := v.Head
		exp := v.Expect

		for res != nil && exp != nil {
			if res.Val != exp.Val {
				t.Errorf("list %s, to delete %d, expect %s, got %s", formatList(head), val, formatList(v.Expect), formatList(v.Head))
				break
			}
			res = res.Next
			exp = exp.Next
		}

		if !(res == nil && exp == nil) {
			t.Errorf("list %s, to delete %d, expect %s, got %s", formatList(head), val, formatList(v.Expect), formatList(v.Head))
		}
	}
}

func copyList(p *ListNode) *ListNode {
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

func formatList(p *ListNode) string {
	s := "["
	for p != nil {
		s += fmt.Sprintf("%d -> ", p.Val)
		p = p.Next
	}
	s += "nil]"
	return s
}
