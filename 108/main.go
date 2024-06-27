package main

import (
	"fmt"
)

func main() {
	root := sortedArrayToBST([]int{0, 1, 2, 3, 4, 5})
	fmt.Println(root.Right.Left.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return findCenterAndReturnRoot(nums)
}

func findCenterAndReturnRoot(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	center := len(nums) / 2
	root := &TreeNode{
		Val:   nums[center],
		Left:  findCenterAndReturnRoot(nums[:center]),
		Right: findCenterAndReturnRoot(nums[center+1:]),
	}
	return root
}
