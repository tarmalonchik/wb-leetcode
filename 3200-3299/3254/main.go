package main

func resultsArray(nums []int, k int) []int {
	if k == 1 {
		return nums
	}

	out := make([]int, 0, len(nums)-k+1)
	pacman := List{maxNum: k}
	newPos := pacman.feed(nums, k)
	out = append(out, pacman.nice())

	for i := newPos; i < len(nums); i++ {
		pacman.eat(nums[i])
		out = append(out, pacman.nice())
	}
	return out
}

type List struct {
	head              *Node
	tail              *Node
	totalNum          int
	maxNum            int
	corruptedElements int
}

func (d *List) nice() int {
	if d.totalNum != d.maxNum {
		return -1
	}
	if d.corruptedElements == 0 && d.totalNum > 0 {
		return d.tail.Number
	}
	return -1
}

func (d *List) feed(arr []int, k int) int {
	for i := 0; i < len(arr); i++ {
		d.eat(arr[i])
		if i+1 == k {
			return i + 1
		}
	}
	return -1
}

func (d *List) eat(number int) {
	if d.totalNum == 0 && d.maxNum > 0 {
		d.pushTail(number)
		return
	}

	lastNum := d.tail.Number
	d.pushTail(number)
	if lastNum+1 != number {
		d.corruptedElements++
	}

	if d.totalNum > d.maxNum {
		d.poop()
	}
}

func (d *List) poop() {
	val, ok := d.removeFront()
	if ok {
		if d.totalNum == 0 {
			return
		}
		if val+1 == d.head.Number {
			return
		}
		d.corruptedElements--
	}
}

func (d *List) removeFront() (int, bool) {
	out := d.head
	if d.head != nil {
		d.head = d.head.Next
		if d.head == nil {
			d.tail = nil
		}
	}
	if out != nil {
		d.totalNum--
		return out.Number, true
	}
	return 0, false
}

func (d *List) pushTail(position int) {
	d.totalNum++
	if d.head == nil && d.tail == nil {
		d.head = &Node{
			Number: position,
		}
		d.tail = d.head
		return
	}
	d.tail.Next = &Node{
		Number: position,
	}
	d.tail = d.tail.Next
}

type Node struct {
	Number int
	Next   *Node
}
