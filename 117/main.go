package main

import (
	"fmt"
)

func main() {
	node := &Node{
		Val: 0,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 1,
				Left: &Node{
					Val: 5,
				},
				Right: &Node{
					Val: 1,
				},
			},
		},
		Right: &Node{
			Val: 4,
			Left: &Node{
				Val: 3,
				Right: &Node{
					Val: 6,
				},
			},
			Right: &Node{
				Val: -1,
				Right: &Node{
					Val: 8,
				},
			},
		},
	}

	connect(node)
	fmt.Println(node.Left.Left.Left.Next.Val)
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	connect(root.Left)
	connect(root.Right)

	if root.Left == nil || root.Right == nil {
		return root
	}

	lCurrent := root.Left
	rCurrent := root.Right
	for {
		if lCurrent == nil {
			break
		}
		lCurrent.Next = rCurrent

		if lCurrent.Right != nil {
			lCurrent = lCurrent.Right
		} else if lCurrent.Left != nil {
			lCurrent = lCurrent.Left
		} else {
			break
		}

		if rCurrent.Left != nil {
			rCurrent = rCurrent.Left
		} else if rCurrent.Right != nil {
			rCurrent = rCurrent.Right
		} else {
			break
		}
	}
	return root
}
