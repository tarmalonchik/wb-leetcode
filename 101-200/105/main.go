package main

import (
	"fmt"
)

func main() {
	preorder := []int{1, 2, 3}
	inorder := []int{1, 3, 2}
	node := buildTree(preorder, inorder)
	fmt.Println(node.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type goBack struct {
	node       *TreeNode
	nextGoBack *goBack
}

func addToGoBack(g *goBack, node *TreeNode) *goBack {
	return &goBack{
		node:       node,
		nextGoBack: g,
	}
}

func getFromGoBack(g *goBack) (*goBack, *TreeNode) {
	return g.nextGoBack, g.node
}

func buildTree(preorder []int, inorder []int) (root *TreeNode) {
	if len(preorder) == 0 {
		return nil
	}

	var (
		inorderPosMap = make(map[int]int, len(inorder))
		back          *goBack
	)

	for i := range inorder {
		inorderPosMap[inorder[i]] = i
	}

	root = &TreeNode{
		Val: preorder[0],
	}

	currentNode := root

	pos := 0

	for {
		if pos >= len(preorder)-1 {
			return
		}

		prevVal := getInorderPrevVal(inorderPosMap, inorder, currentNode.Val)
		//fmt.Println(*prevVal, currentNode.Val)
		postVal := getInorderPostVal(inorderPosMap, inorder, currentNode.Val)
		//fmt.Println(*postVal, currentNode.Val)

		if prevVal != nil && preorder[pos+1] == *prevVal {
			currentNode.Left = &TreeNode{
				Val: preorder[pos+1],
			}
			back = addToGoBack(back, currentNode)
			currentNode = currentNode.Left
		} else if postVal != nil && preorder[pos+1] == *postVal {
			currentNode.Right = &TreeNode{
				Val: preorder[pos+1],
			}
			currentNode = currentNode.Right
		} else {
			if back == nil {
				if prevVal == nil {
					currentNode.Right = &TreeNode{
						Val: preorder[pos+1],
					}
					currentNode = currentNode.Right
					pos++
					continue
				}
				currentNode.Left = &TreeNode{
					Val: preorder[pos+1],
				}
				currentNode = currentNode.Left
				pos++
				continue
			}
			back, currentNode = getFromGoBack(back)
			currentNode.Right = &TreeNode{
				Val: preorder[pos+1],
			}
			currentNode = currentNode.Right
		}
		pos++
	}
}

func getInorderPrevVal(inorderPosMap map[int]int, inorder []int, val int) *int {
	pos, ok := inorderPosMap[val]
	if !ok {
		return nil
	}

	pos = pos - 1

	if pos < 0 {
		return nil
	}

	return &inorder[pos]
}

func getInorderPostVal(inorderPosMap map[int]int, inorder []int, val int) *int {
	pos, ok := inorderPosMap[val]
	if !ok {
		return nil
	}

	pos = pos + 1

	if pos >= len(inorder) {
		return nil
	}

	return &inorder[pos]
}
