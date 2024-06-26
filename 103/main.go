package main

import (
	"fmt"
)

func main() {
	//node :=
	fmt.Println(zigzagLevelOrder(nil))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type chain struct {
	node *TreeNode
	next *chain
}

func addToChain(g *chain, node *TreeNode) *chain {
	return &chain{
		node: node,
		next: g,
	}
}

func getFromChain(g *chain) (*chain, *TreeNode) {
	return g.next, g.node
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var (
		resp               [][]int
		level              int
		chainOne, chainTwo *chain
	)

	chainOne = &chain{
		node: root,
	}
	for {
		var node *TreeNode

		chainOne, node = getFromChain(chainOne)

		resp = addToResp(resp, level, node.Val)

		if node.Right == nil || node.Left == nil {
			if node.Right != nil {
				chainTwo = addToChain(chainTwo, node.Right)
			} else if node.Left != nil {
				chainTwo = addToChain(chainTwo, node.Left)
			}
		} else if node.Right != nil && node.Left != nil {
			if level%2 == 0 {
				chainTwo = addToChain(chainTwo, node.Left)
				chainTwo = addToChain(chainTwo, node.Right)
			} else {
				chainTwo = addToChain(chainTwo, node.Right)
				chainTwo = addToChain(chainTwo, node.Left)
			}
		}

		if chainOne == nil {
			if chainTwo == nil {
				break
			}
			chainOne = chainTwo
			chainTwo = nil
			level++
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
