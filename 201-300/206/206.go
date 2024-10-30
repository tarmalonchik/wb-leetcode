package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	first, last := recursive(head)
	if last != nil {
		last.Next = nil
	}
	return first
}

func recursive(item *ListNode) (first, last *ListNode) {
	if item.Next == nil {
		return item, item.Next
	}
	first, last = recursive(item.Next)
	if last == nil {
		last = item
		first.Next = last
	} else {
		last.Next = item
		last = last.Next
	}
	return
}
