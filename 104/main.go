package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type goBack struct {
	node       *TreeNode
	nextGoBack *goBack
	level      int
}

func addToGoBack(g *goBack, node *TreeNode, level int) *goBack {
	return &goBack{
		node:       node,
		nextGoBack: g,
		level:      level,
	}
}

func getFromGoBack(g *goBack) (*goBack, *TreeNode, int) {
	return g.nextGoBack, g.node, g.level
}

func maxDepth(root *TreeNode) (maxDepth int) {
	if root == nil {
		return 0
	}

	var (
		level int
		back  *goBack
	)

	for {
		if root.Left != nil && root.Right != nil {
			back = addToGoBack(back, root.Right, level+1)
			root = root.Left
			level++
			if level > maxDepth {
				maxDepth = level
			}
			continue
		}

		if root.Left == nil && root.Right == nil {
			if back == nil {
				break
			}
			back, root, level = getFromGoBack(back)
			continue
		}

		level++
		if level > maxDepth {
			maxDepth = level
		}
		if root.Left != nil {
			root = root.Left
			continue
		}
		if root.Right != nil {
			root = root.Right
			continue
		}
	}
	return maxDepth + 1
}
