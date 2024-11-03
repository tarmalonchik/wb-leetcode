package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}

	originHead := head

	counter := 1

	for {
		if head.Next == nil {
			break
		}
		if left == 1 {
			originHead = subReverse(head, counter, right)
			break
		}
		if counter == left-1 {
			newSubHead := subReverse(head.Next, counter+1, right)
			head.Next = newSubHead
			break
		}
		counter++
		head = head.Next
	}
	return originHead
}

func subReverse(head *ListNode, counter int, finishPos int) (start *ListNode) {
	startNode := head
	currentNode := head

	var prevNode *ListNode

	for {
		if counter == finishPos {
			startNode.Next = currentNode.Next
			currentNode.Next = prevNode
			return currentNode
		}
		temp := currentNode.Next
		currentNode.Next = prevNode
		prevNode = currentNode
		currentNode = temp
		counter++
	}
}
