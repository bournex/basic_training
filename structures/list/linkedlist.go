package list

import "fmt"

type ListNode struct {
	Next  *ListNode
	Value interface{}
}

type LinkedList struct {
	Head *ListNode
	Tail *ListNode
}

func MakeLinkedList() *LinkedList {
	return new(LinkedList)
}

func (sll *LinkedList) Print() {
	head := sll.Head

	for head != nil {
		fmt.Println(head.Value)
		head = head.Next
	}
}

func (sll *LinkedList) Append(node *ListNode) {
	if sll.Head == nil && sll.Tail == nil {
		// 空链表
		sll.Head = node
		sll.Tail = node
	} else {
		sll.Tail.Next = node
		sll.Tail = node
	}
}

func (sll *LinkedList) RemoveHead() {
	if sll.Head != sll.Tail {
		// 空链表
		sll.Head = sll.Head.Next
	}
}

func (sll *LinkedList) RemoveTail() {
	p := sll.Head
	if p == nil {
		return
	}

	if p.Next == nil {
		sll.Head = nil
		sll.Tail = nil
	}

	for p.Next != sll.Tail {
		p = p.Next
	}

	sll.Tail = p
	p.Next = nil
}
