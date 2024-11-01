package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type List struct {
	Next           *List
	Value          *TreeNode
	ProcessedCount uint8
}

func (l *List) GetChildrenCount() int {
	if l.Value.Right != nil && l.Value.Left != nil {
		return 2
	}
	if l.Value.Right == nil && l.Value.Left == nil {
		return 0
	}
	return 1
}

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}

	list := addToList(nil, root)

	for {
		if root.Left == nil {
			break
		}
		list.ProcessedCount = 1
		list = addToList(list, root.Left)
		root = root.Left
	}

	for {
		if list == nil {
			break
		}

		if list.ProcessedCount == 0 {
			childrenCount := list.GetChildrenCount()

			if childrenCount == 0 {
				k--
				if k == 0 {
					return list.Value.Val
				}
				list = removeFromList(list)
				continue
			}

			list.ProcessedCount++
			if list.Value.Left != nil {
				list = addToList(list, list.Value.Left)
			} else {
				k--
				if k == 0 {
					return list.Value.Val
				}
				list = addToList(list, list.Value.Right)
			}
			continue

		} else if list.ProcessedCount == 1 {
			childrenCount := list.GetChildrenCount()

			if childrenCount == 1 {
				if list.Value.Left != nil {
					k--
					if k == 0 {
						return list.Value.Val
					}
					list = removeFromList(list)
					continue
				}
				list = removeFromList(list)
				continue
			}
			k--
			if k == 0 {
				return list.Value.Val
			}
			node := list.Value.Right
			list = removeFromList(list)
			list = addToList(list, node)
		}
	}
	return 0
}

func removeFromList(list *List) *List {
	if list == nil {
		return nil
	}
	return list.Next
}

func addToList(list *List, node *TreeNode) *List {
	if list == nil {
		return &List{
			Value: node,
			Next:  nil,
		}
	}
	return &List{
		Value: node,
		Next:  list,
	}
}
