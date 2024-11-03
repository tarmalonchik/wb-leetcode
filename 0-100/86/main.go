package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}

	var lessHead *ListNode
	var lessTail *ListNode
	var grossHead *ListNode
	var grossTail *ListNode

	currentNode := head
	for {
		if currentNode.Val < x {
			if lessHead == nil {
				lessHead = currentNode
				lessTail = currentNode
			} else {
				lessTail.Next = currentNode
				lessTail = lessTail.Next
			}
		} else {
			if grossHead == nil {
				grossHead = currentNode
				grossTail = currentNode
			} else {
				grossTail.Next = currentNode
				grossTail = grossTail.Next
			}
		}
		currentNode = currentNode.Next
		if currentNode == nil {
			break
		}
	}

	if lessHead == nil {
		return grossHead
	}
	lessTail.Next = grossHead
	if grossTail != nil {
		grossTail.Next = nil
	}
	return lessHead
}
