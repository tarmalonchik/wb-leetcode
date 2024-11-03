package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}

	listLen := 0
	var newHead = head
	var oldTail *ListNode

	for {
		if newHead == nil {
			break
		}
		listLen++
		if newHead.Next == nil {
			oldTail = newHead
		}
		newHead = newHead.Next
	}

	k = k % listLen

	counter := 1
	newHead = head

	for {
		if counter == listLen-k {
			oldTail.Next = head
			head = newHead.Next
			newHead.Next = nil
			return head
		}
		counter++
		newHead = newHead.Next
	}
}
