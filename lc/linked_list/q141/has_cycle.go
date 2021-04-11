package q141

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

// 思路
//	O(N)空间复杂度解法用哈希表保存节点
//	O(1)空间复杂度解法用快慢指针
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			return false
		}
		if slow == fast {
			return true
		}
	}
	return false
}
