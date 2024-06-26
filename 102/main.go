package main

import (
	"fmt"
)

func main() {
	node := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 88,
				Left: &TreeNode{
					Val: 11,
				},
			},
		},
	}
	fmt.Println(levelOrder(node))
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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var (
		level int
		resp  [][]int
		back  *goBack
	)

	if root == nil {
		return nil
	}

	for {
		resp = addToResp(resp, level, root.Val)

		if root.Left != nil && root.Right != nil {
			back = addToGoBack(back, root.Right, level+1)
			root = root.Left
			level++
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
		if root.Left != nil {
			root = root.Left
			continue
		}
		if root.Right != nil {
			root = root.Right
			continue
		}
	}
	return resp
}

func addToResp(resp [][]int, level, val int) [][]int {
	if len(resp) == level {
		resp = append(resp, []int{val})
		return resp
	}
	resp[level] = append(resp[level], val)
	return resp
}
