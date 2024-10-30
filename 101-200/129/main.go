package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type goBack struct {
	node       *TreeNode
	nextGoBack *goBack
	value      int
}

func addToGoBack(g *goBack, node *TreeNode, val int) *goBack {
	return &goBack{
		node:       node,
		nextGoBack: g,
		value:      val,
	}
}

func getFromGoBack(g *goBack) (*goBack, *TreeNode, int) {
	return g.nextGoBack, g.node, g.value
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var back *goBack
	var resp int
	currentValue := root.Val

	for {
		if root.Left == nil && root.Right == nil {
			resp += currentValue
			if back == nil {
				return resp
			}
			back, root, currentValue = getFromGoBack(back)
			root = root.Right
			currentValue = currentValue * 10
			currentValue += root.Val
			continue
		}

		if root.Left != nil && root.Right != nil {
			back = addToGoBack(back, root, currentValue)
			root = root.Left
			currentValue = currentValue * 10
			currentValue += root.Val
			continue
		}

		if root.Left != nil {
			root = root.Left
			currentValue = currentValue * 10
			currentValue += root.Val
			continue
		}
		if root.Right != nil {
			root = root.Right
			currentValue = currentValue * 10
			currentValue += root.Val
			continue
		}
	}
}
