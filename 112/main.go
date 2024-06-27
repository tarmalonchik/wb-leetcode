package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type goBack struct {
	node       *TreeNode
	nextGoBack *goBack
	prevSum    int
}

func addToGoBack(g *goBack, node *TreeNode, prevSum int) *goBack {
	return &goBack{
		node:       node,
		nextGoBack: g,
		prevSum:    prevSum,
	}
}

func getFromGoBack(g *goBack) (*goBack, *TreeNode, int) {
	return g.nextGoBack, g.node, g.prevSum
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	var back *goBack
	var localSum int

	if root == nil {
		return false
	}

	localSum += root.Val
	for {
		if root.Left == nil && root.Right == nil {
			if targetSum == localSum {
				return true
			}
			if back == nil {
				return false
			}
			back, root, localSum = getFromGoBack(back)
			root = root.Right
			localSum += root.Val
			continue
		}

		if root.Left != nil && root.Right != nil {
			back = addToGoBack(back, root, localSum)
			root = root.Left
			localSum += root.Val
			continue
		}

		if root.Left != nil {
			root = root.Left
			localSum += root.Val
			continue
		}

		if root.Right != nil {
			root = root.Right
			localSum += root.Val
			continue
		}
	}
}
