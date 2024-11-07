package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	var newHead *ListNode
	var prevTail *ListNode

	for {
		if head == nil || head.Next == nil {
			if newHead == nil {
				newHead = head
			}
			break
		}

		if newHead == nil {
			newHead = head.Next
		}

		nextHead := head.Next.Next
		prevTail = justSwap(head, prevTail)
		head = nextHead
	}
	return newHead
}

func justSwap(current, prev *ListNode) (tail *ListNode) {
	if current == nil || current.Next == nil {
		return current
	}

	if prev != nil {
		prev.Next = current.Next
	}
	nextNext := current.Next.Next
	current.Next.Next = current
	current.Next = nextNext
	return current
}
