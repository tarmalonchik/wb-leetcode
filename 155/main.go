package main

type MinStack struct {
	Tail *StackNode
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	minVal := 0

	if this.Tail == nil {
		minVal = val
	} else {
		if this.Tail.MinValue < val {
			minVal = this.Tail.MinValue
		} else {
			minVal = val
		}
	}

	newItem := &StackNode{
		Next:     this.Tail,
		Val:      val,
		MinValue: minVal,
	}

	this.Tail = newItem
}

func (this *MinStack) Pop() {
	if this.Tail != nil {
		this.Tail = this.Tail.Next
	}
}

func (this *MinStack) Top() int {
	if this.Tail == nil {
		return 0
	}
	return this.Tail.Val
}

func (this *MinStack) GetMin() int {
	if this.Tail == nil {
		return 0
	}
	return this.Tail.MinValue
}

type StackNode struct {
	Val      int
	Next     *StackNode
	MinValue int
}
