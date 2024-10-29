package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var currentPos = head
	var prevPosition *ListNode

	for {
		if currentPos == nil {
			if prevPosition == nil {
				return nil
			}
			break
		}

		if currentPos.Next != nil && currentPos.Next.Next != nil {
			if currentPos.Val == currentPos.Next.Val && currentPos.Val == currentPos.Next.Next.Val {
				currentPos = currentPos.Next
				continue
			}
		}

		if currentPos.Next != nil {
			if currentPos.Val == currentPos.Next.Val {
				currentPos = currentPos.Next.Next
				if currentPos == nil && prevPosition != nil {
					prevPosition.Next = nil
				}
				continue
			}
		}

		if prevPosition == nil {
			prevPosition = currentPos
			head = prevPosition
		} else {
			prevPosition.Next = currentPos
			prevPosition = prevPosition.Next
		}
		currentPos = currentPos.Next
	}
	return head
}
