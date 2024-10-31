package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow := head
	fast := head

	for {
		if fast == nil || fast.Next == nil {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return false
		}
	}
}
