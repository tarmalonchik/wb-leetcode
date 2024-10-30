package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	n++
	var lag *ListNode
	var lead *ListNode
	lead = head
	for {
		if lag == nil && n >= 0 {
			if n == 0 {
				lag = head
			} else {
				n--
			}
		} else {
			lag = lag.Next
		}
		if lead != nil {
			lead = lead.Next
		} else {
			if n > 0 {
				return head
			}
			if n == 0 && lag != nil {
				remove(lag)
				return head
			}
			return head.Next
		}
	}
}

func remove(preNodeToRemove *ListNode) {
	if preNodeToRemove != nil {
		if preNodeToRemove.Next != nil {
			preNodeToRemove.Next = preNodeToRemove.Next.Next
		} else {
			preNodeToRemove = nil
		}
	}
}
