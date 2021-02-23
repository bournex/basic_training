package list

// 翻转单链表

func (sll *LinkedList) Reverse1() {
	var (
		prev *ListNode
		curr *ListNode
	)

	curr = sll.Head
	sll.Tail = sll.Head
	for curr != nil {
		sll.Head = curr.Next
		curr.Next = prev
		prev = curr
		curr = sll.Head
	}
	sll.Head = prev
}
