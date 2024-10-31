package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type List struct {
	Node      *TreeNode
	Next      *List
	Level     int
	Processed bool
}

func AddToList(list *List, node *TreeNode, level int) *List {
	if node == nil {
		return list
	}
	if list == nil {
		return &List{
			Node:  node,
			Level: level,
		}
	}
	return &List{
		Node:      node,
		Next:      list,
		Level:     level,
		Processed: false,
	}
}

func RemoveFromList(list *List) *List {
	if list == nil {
		return nil
	}
	return list.Next
}

func countNodes(root *TreeNode) int {
	var number int

	if root == nil {
		return 0
	}

	list := &List{
		Node:      root,
		Next:      nil,
		Level:     0,
		Processed: false,
	}

	for {
		if list.Node.Left != nil && list.Node.Right != nil {
			if list.Processed {
				right := list.Node.Right
				level := list.Level
				list = RemoveFromList(list)
				list = AddToList(list, right, level+1)
				continue
			}
			number++
			list.Processed = true
			list = AddToList(list, list.Node.Left, list.Level+1)
			continue
		} else if list.Node.Left != nil {
			number++
			left := list.Node.Left
			level := list.Level
			list = RemoveFromList(list)
			list = AddToList(list, left, level+1)
			continue
		} else if list.Node.Right != nil {
			number++
			right := list.Node.Right
			level := list.Level
			list = RemoveFromList(list)
			list = AddToList(list, right, level+1)
			continue
		}

		number++

		if list.Next == nil {
			break
		}
		list = list.Next
	}

	return number
}
