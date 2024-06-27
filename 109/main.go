package main

import (
	"fmt"
)

func main() {
	some := []int{}
	root := generateListNode(some)
	tree := sortedListToBST(root)

	fmt.Println(tree.Left.Right.Val)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateListNode(in []int) (resp *ListNode) {
	for i := len(in) - 1; i >= 0; i-- {
		if resp == nil {
			resp = &ListNode{
				Val: in[i],
			}
			continue
		}

		resp = &ListNode{
			Val:  in[i],
			Next: resp,
		}
	}
	return resp
}

func countListNode(in *ListNode) (resp int) {
	currentNode := in
	for {
		if currentNode == nil {
			return
		}
		currentNode = currentNode.Next
		resp++
	}
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	number := countListNode(head)
	node, _ := process(head, number)
	return node
}

func process(head *ListNode, count int) (*TreeNode, *ListNode) {
	if count == 1 {
		return &TreeNode{
			Val: head.Val,
		}, head.Next
	}
	if count == 2 {
		return &TreeNode{
			Val: head.Val,
			Right: &TreeNode{
				Val: head.Next.Val,
			},
		}, head.Next.Next
	}
	if count == 3 {
		return &TreeNode{
			Val: head.Next.Val,
			Right: &TreeNode{
				Val: head.Next.Next.Val,
			},
			Left: &TreeNode{
				Val: head.Val,
			},
		}, head.Next.Next.Next
	}

	resp := &TreeNode{}
	leftCount := count / 2
	if count%2 == 0 {
		leftCount--
	}
	resp.Left, head = process(head, leftCount)
	resp.Val = head.Val
	head = head.Next
	resp.Right, head = process(head, count/2)
	return resp, head
}
