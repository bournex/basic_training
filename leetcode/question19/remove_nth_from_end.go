package question19

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

// 虚拟头结点是很好用滴
// 思路就是双指针移动，两个指针始终保持同等距离，注意前面的指针开始时不是指向链表头，而是链表虚拟头，这样待删除节点就在prev.Next，否则还需要再遍历才能删除

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	prev := &ListNode{0, head}
	tail := head
	head = prev
	for ; tail != nil && n > 0; n-- {
		tail = tail.Next
	}

	for tail != nil {
		prev = prev.Next
		tail = tail.Next
	}

	if prev.Next != nil {
		prev.Next = prev.Next.Next
	}

	return head.Next
}

/*
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	tail, prev := head, head
	for ; tail != nil && n > 0; n-- {
		tail = tail.Next
	}

	for tail != nil && tail.Next != nil {
		prev = prev.Next
		tail = tail.Next
	}

	if prev != nil && prev.Next != nil {
		prev.Next = prev.Next.Next
	}
	return head
}
*/
