package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	floatingNode := head
	for {
		if floatingNode.Next == nil {
			break
		}
		if floatingNode.Next.Val == floatingNode.Val {
			floatingNode.Next = floatingNode.Next.Next
		} else {
			floatingNode = floatingNode.Next
		}
	}
	return head
}
