package main

func main() {
	//isSymmetric()
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return compare(root.Left, root.Right)
}

func compare(lNode, rNode *TreeNode) bool {
	if lNode == nil && rNode == nil {
		return true
	}
	if lNode == nil || rNode == nil {
		return false
	}

	if lNode.Val != rNode.Val {
		return false
	}

	if !compare(lNode.Left, rNode.Right) {
		return false
	}
	return compare(lNode.Right, rNode.Left)
}
