package m86

import ll "github.com/bournex/basic_training/lc/linked_list"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ll.ListNode, x int) *ll.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	return nil
}
