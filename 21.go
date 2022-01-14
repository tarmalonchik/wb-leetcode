package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		firstVal *ListNode
		preVal   *ListNode
	)

	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		firstVal = l1
	} else {
		firstVal = l2
		l2 = l1
		l1 = firstVal
	}

	for {
		if l1.Next == nil {
			l1.Next = l2
			return firstVal
		}
		if l1.Next.Val > l2.Val {
			preVal = l1.Next
			l1.Next = l2
			l2 = preVal
		}
		l1 = l1.Next
	}
}
