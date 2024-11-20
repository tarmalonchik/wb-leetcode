package main

import (
	"fmt"
)

func main() {
	trees := generateTrees(4)
	fmt.Println("response", len(trees))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeNodeWithTail struct {
	TreeNodeRoot *TreeNode
	TreeNodeTail *TreeNode
}

func generateTrees(n int) (out []*TreeNode) {
	if n == 1 {
		return []*TreeNode{
			{
				Val: 1,
			},
		}
	}

	for i := 1; i <= n; i++ {
		var newTrees []*TreeNodeWithTail
		if i == 1 {
			newTrees = generateTreeRecursive(1+1, n, nil, nil, i)
		} else if i == n {
			newTrees = generateTreeRecursive(1, n-1, nil, nil, i)
		} else {
			newTreeLeft := generateTreeRecursive(1, i-1, nil, nil, i)
			for j := range newTreeLeft {
				newData := generateTreeRecursive(i+1, n, newTreeLeft[j].TreeNodeRoot, newTreeLeft[j].TreeNodeTail, 0)
				newTrees = append(newTrees, newData...)
			}
		}
		for j := range newTrees {
			out = append(out, newTrees[j].TreeNodeRoot)
		}
	}
	return out
}

func generateTreeRecursive(first, last int, parentRoot, parentTail *TreeNode, baseVal int) (out []*TreeNodeWithTail) {
	if baseVal != 0 {
		parentRoot = &TreeNode{
			Val: baseVal,
		}
		parentTail = parentRoot
	}

	if first == last {
		if parentTail.Val < first {
			parentTail.Right = &TreeNode{
				Val: first,
			}
		} else {
			parentTail.Left = &TreeNode{
				Val: first,
			}
		}
		out = append(out, &TreeNodeWithTail{
			TreeNodeRoot: parentRoot,
			TreeNodeTail: parentTail,
		})
	}
	if last-first == 1 {
		if parentTail.Val < first {
			newHead, newTail := copyTree(parentRoot, parentTail.Val)

			parentTail.Right = &TreeNode{
				Val: first,
				Right: &TreeNode{
					Val: last,
				},
			}
			newTail.Right = &TreeNode{
				Val: last,
				Left: &TreeNode{
					Val: first,
				},
			}
			out = append(out, &TreeNodeWithTail{
				TreeNodeRoot: parentRoot,
				TreeNodeTail: parentTail,
			})
			out = append(out, &TreeNodeWithTail{
				TreeNodeRoot: newHead,
				TreeNodeTail: newTail,
			})
		} else if parentTail.Val > last {
			newHead, newTail := copyTree(parentRoot, parentTail.Val)

			parentTail.Left = &TreeNode{
				Val: last,
				Left: &TreeNode{
					Val: first,
				},
			}
			newTail.Left = &TreeNode{
				Val: first,
				Right: &TreeNode{
					Val: last,
				},
			}
			out = append(out, &TreeNodeWithTail{
				TreeNodeRoot: parentRoot,
				TreeNodeTail: parentTail,
			})
			out = append(out, &TreeNodeWithTail{
				TreeNodeRoot: newHead,
				TreeNodeTail: newTail,
			})
		}
		return out
	}

	for i := first; i <= last; i++ {
		if i == first {
			newHead, newTail := copyTree(parentRoot, parentTail.Val)
			if i < newTail.Val {
				newTail.Left = &TreeNode{
					Val: i,
				}
				newTrees := generateTreeRecursive(first+1, last, newHead, newTail.Left, 0)
				out = append(out, newTrees...)
			} else {
				newTail.Right = &TreeNode{
					Val: i,
				}
				newTrees := generateTreeRecursive(first+1, last, newHead, newTail.Right, 0)
				out = append(out, newTrees...)
			}
		} else if i == last {
			newHead, newTail := copyTree(parentRoot, parentTail.Val)

			if i < newTail.Val {
				newTail.Left = &TreeNode{
					Val: i,
				}
				newTrees := generateTreeRecursive(first, last-1, newHead, newTail.Left, 0)
				out = append(out, newTrees...)
			} else {
				newTail.Right = &TreeNode{
					Val: i,
				}
				newTrees := generateTreeRecursive(first, last-1, newHead, newTail.Right, 0)
				out = append(out, newTrees...)
			}
		} else {
			newHead, newTail := copyTree(parentRoot, parentTail.Val)
			if i < newTail.Val {
				newTail.Left = &TreeNode{
					Val: i,
				}
				left := generateTreeRecursive(first, i-1, newHead, newTail.Left, 0)
				for j := range left {
					out = append(out, generateTreeRecursive(i+1, last, left[j].TreeNodeRoot, left[j].TreeNodeTail, 0)...)
				}
			} else {
				newTail.Right = &TreeNode{
					Val: i,
				}
				left := generateTreeRecursive(first, i-1, newHead, newTail.Right, 0)
				for j := range left {
					out = append(out, generateTreeRecursive(i+1, last, left[j].TreeNodeRoot, left[j].TreeNodeTail, 0)...)
				}
			}
		}
	}
	return out
}

func copyTree(root *TreeNode, rootTailVal int) (outHead, outTail *TreeNode) {
	if root == nil {
		return nil, nil
	}

	outHead = &TreeNode{
		Val: root.Val,
	}

	left, tailCandidate := copyTree(root.Left, rootTailVal)
	if tailCandidate != nil {
		outTail = tailCandidate
	}

	right, tailCandidate2 := copyTree(root.Right, rootTailVal)
	if tailCandidate2 != nil {
		outTail = tailCandidate2
	}

	outHead.Left = left
	outHead.Right = right

	if outHead.Val == rootTailVal {
		outTail = outHead
	}

	return outHead, outTail
}
