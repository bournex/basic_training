package h25

import ll "github.com/bournex/basic_training/lc/linked_list"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 思路
//	没有花哨的技巧，主要是细节处理，还是训练对链表操作的熟练度
//	一般性技巧：对于需要对某一段子链表操作的情形，可以建立虚拟头结点

func reverseKGroup(head *ll.ListNode, k int) *ll.ListNode {
	dummy := &ll.ListNode{Next: head}

	reverse := func(p *ll.ListNode) *ll.ListNode {
		var shorthead *ll.ListNode
		shortcurr := p
		for shortcurr != nil {
			shortnext := shortcurr.Next
			shortcurr.Next = shorthead
			shorthead = shortcurr
			shortcurr = shortnext
		}
		return shorthead
	}

	subhead, subtail := dummy, dummy

	for subtail.Next != nil {
		i := 0
		for i < k && subtail.Next != nil {
			subtail = subtail.Next
			i++
		}

		if subtail.Next != nil || i == k {
			temp := subtail.Next
			subtail.Next = nil
			newtail := subhead.Next
			reverse(subhead.Next)
			subhead.Next = subtail
			subhead = newtail
			subhead.Next = temp
			subtail = subhead
		}
	}

	return dummy.Next
}
