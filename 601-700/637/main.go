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

func averageOfLevels(root *TreeNode) []float64 {
	var levelsCounter [][]int

	if root == nil {
		return nil
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
			levelsCounter = addToLevelsCounter(levelsCounter, list.Node.Val, list.Level)
			list.Processed = true
			list = AddToList(list, list.Node.Left, list.Level+1)
			continue
		} else if list.Node.Left != nil {
			levelsCounter = addToLevelsCounter(levelsCounter, list.Node.Val, list.Level)
			left := list.Node.Left
			level := list.Level
			list = RemoveFromList(list)
			list = AddToList(list, left, level+1)
			continue
		} else if list.Node.Right != nil {
			levelsCounter = addToLevelsCounter(levelsCounter, list.Node.Val, list.Level)
			right := list.Node.Right
			level := list.Level
			list = RemoveFromList(list)
			list = AddToList(list, right, level+1)
			continue
		}

		levelsCounter = addToLevelsCounter(levelsCounter, list.Node.Val, list.Level)

		if list.Next == nil {
			break
		}
		list = list.Next
	}

	out := make([]float64, len(levelsCounter))

	for i := range levelsCounter {
		for j := range levelsCounter[i] {
			out[i] += float64(levelsCounter[i][j])
		}
		out[i] = out[i] / float64(len(levelsCounter[i]))
	}

	return out
}

func addToLevelsCounter(levelsCounter [][]int, val, level int) [][]int {
	if len(levelsCounter) == level {
		levelsCounter = append(levelsCounter, []int{})
	} else if len(levelsCounter) < level {
		panic("invalid flow")
	}

	levelsCounter[level] = append(levelsCounter[level], val)
	return levelsCounter
}
