package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil || root.Left == nil {
		return root
	}

	lCurrent := root.Left
	rCurrent := root.Right
	for {
		if lCurrent == nil {
			break
		}
		lCurrent.Next = rCurrent
		lCurrent = lCurrent.Right
		rCurrent = rCurrent.Left
	}

	connect(root.Left)
	connect(root.Right)
	return root
}
