package q2

import "testing"

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}     // 342
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}     // 465
	expect := &ListNode{7, &ListNode{0, &ListNode{8, nil}}} // 807

	res := addTwoNumbers(l1, l2)

	for res != nil {
		if res.Val != expect.Val {
			t.Errorf("expect %d, got %d", expect.Val, res.Val)
		}
		res = res.Next
		expect = expect.Next
	}

	if expect != nil {
		t.Errorf("length not the same")
	}
}
