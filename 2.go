package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	flag := 0
	if l1 != nil || l2 != nil {
		firstNode := &ListNode{
			Val:  0,
			Next: nil,
		}
		firstNode.Val += l1.Val
		firstNode.Val += l2.Val
		if firstNode.Val > 9 {
			flag = 1
			firstNode.Val %= 10
		}
		if flag == 1 || l1.Next != nil || l2.Next != nil {
			curr := &ListNode{}
			firstNode.Next = curr
			var curr1 *ListNode
			var curr2 *ListNode
			if l1.Next != nil {
				curr1 = l1.Next
			}
			if l2.Next != nil {
				curr2 = l2.Next
			}
			for {
				if curr1 == nil && curr2 == nil && flag == 0 {
					return firstNode
				}
				curr.Val = flag
				flag = 0
				if curr1 != nil {
					curr.Val += curr1.Val
					curr1 = curr1.Next
				}
				if curr2 != nil {
					curr.Val += curr2.Val
					curr2 = curr2.Next
				}
				if curr.Val > 9 {
					flag = 1
					curr.Val %= 10
				}
				
				if flag == 1 || curr1 != nil || curr2 != nil {
					curr.Next = &ListNode{}
				}
				curr = curr.Next
			}
		}
		return firstNode
	}
	return &ListNode{}
}
