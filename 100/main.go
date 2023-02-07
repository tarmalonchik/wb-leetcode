package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == q {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	if p.Right != q.Right {
		if !isSameTree(p.Right, q.Right) {
			return false
		}
	}
	if p.Left != q.Left {
		if !isSameTree(p.Left, q.Left) {
			return false
		}
	}
	return true
}
