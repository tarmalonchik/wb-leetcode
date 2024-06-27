package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type goBack struct {
	node       *TreeNode
	nextGoBack *goBack
	prevArr    []int
	prevSum    int
}

func addToGoBack(g *goBack, node *TreeNode, prevArr []int, prevSum int) *goBack {
	return &goBack{
		node:       node,
		nextGoBack: g,
		prevArr:    prevArr,
		prevSum:    prevSum,
	}
}

func getFromGoBack(g *goBack) (*goBack, *TreeNode, []int, int) {
	return g.nextGoBack, g.node, g.prevArr, g.prevSum
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	var back *goBack
	var resp [][]int
	var localResp []int
	var localSum int

	if root == nil {
		return nil
	}

	localResp = append(localResp, root.Val)
	localSum += root.Val

	for {
		if root.Left == nil && root.Right == nil {
			if targetSum == localSum {
				respAdd := make([]int, len(localResp))
				copy(respAdd, localResp)
				resp = append(resp, respAdd)
			}
			if back == nil {
				return resp
			}
			back, root, localResp, localSum = getFromGoBack(back)
			root = root.Right
			localResp = append(localResp, root.Val)
			localSum += root.Val
			continue
		}

		if root.Left != nil && root.Right != nil {
			back = addToGoBack(back, root, localResp, localSum)
			root = root.Left
			localResp = append(localResp, root.Val)
			localSum += root.Val
			continue
		}

		if root.Left != nil {
			root = root.Left
			localResp = append(localResp, root.Val)
			localSum += root.Val
			continue
		}

		if root.Right != nil {
			fmt.Println("here")

			root = root.Right
			localResp = append(localResp, root.Val)
			localSum += root.Val
			continue
		}
	}
}
