package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	var minVal = -1
	var currentValue = 0

	if root == nil {
		return 0
	}

	var backWay *List

	for {
		backWay = Add(backWay, root)
		currentValue = root.Val
		if root.Left == nil {
			break
		}
		root = root.Left
	}

	for {
		if backWay.TreeNode.Left != nil && backWay.TreeNode.Left.Val > currentValue {
			backWay = Add(backWay, backWay.TreeNode.Left)

			if backWay.TreeNode.Left == nil {
				minVal = setMin(minVal, backWay.TreeNode.Val, currentValue)
				if currentValue < backWay.TreeNode.Val {
					currentValue = backWay.TreeNode.Val
				}
			}
			continue
		}

		if backWay.TreeNode.Right != nil && backWay.TreeNode.Right.Val > currentValue {
			backWay = Add(backWay, backWay.TreeNode.Right)
			if backWay.TreeNode.Left == nil {
				minVal = setMin(minVal, backWay.TreeNode.Val, currentValue)
				if currentValue < backWay.TreeNode.Val {
					currentValue = backWay.TreeNode.Val
				}
			}
			continue
		}

		if backWay.Next == nil {
			break
		}

		minVal = setMin(minVal, currentValue, backWay.Next.TreeNode.Val)

		backWay = GetNext(backWay)

		if currentValue < backWay.TreeNode.Val {
			currentValue = backWay.TreeNode.Val
		}
	}

	return minVal
}

func setMin(origValue, a, b int) int {
	if origValue < 0 {
		return abs(a - b)
	}
	if origValue > abs(a-b) {
		origValue = abs(a - b)
	}
	return origValue
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

type List struct {
	Next     *List
	TreeNode *TreeNode
}

func GetNext(in *List) *List {
	if in == nil {
		return nil
	}
	return in.Next
}

func Add(in *List, node *TreeNode) *List {
	if in == nil {
		return &List{
			Next:     nil,
			TreeNode: node,
		}
	}
	return &List{
		Next:     in,
		TreeNode: node,
	}
}
