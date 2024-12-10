package main

func maximumLength(s string) int {
	if len(s) < 3 {
		return -1
	}

	minSize := 1
	maxSize := len(s) - 2

	for {
		if maxSize-minSize <= 1 {
			if check(s, maxSize) {
				return maxSize
			}
			if check(s, minSize) {
				return minSize
			}
			return -1
		}
		center := minSize + (maxSize-minSize)/2
		if check(s, center) {
			minSize = center
		} else {
			maxSize = center
		}
	}
}

func check(s string, size int) bool {
	pac := newPacman(size)
	startPos := pac.feed(s)
	container := make(map[byte]int)
	for {
		valid, symbol := pac.nice()
		if valid {
			if val, ok := container[symbol]; ok {
				container[symbol] = val + 1
				if val+1 >= 3 {
					//fmt.Println(string(symbol))
					return true
				}
			} else {
				container[symbol] = 1
			}
		}
		startPos++
		if startPos > len(s)-1 {
			break
		}
		pac.eat(s[startPos])
	}
	return false
}

func newPacman(capacity int) pacman {
	return pacman{
		capacity: capacity,
		size:     0,
	}
}

type pacman struct {
	head         *Node
	tail         *Node
	capacity     int
	size         int
	diffElements int
}

func (d *pacman) nice() (bool, byte) {
	if d.size != d.capacity {
		return false, 0
	}
	if d.diffElements == 1 {
		return true, d.head.Symbol
	}
	return false, 0
}

func (d *pacman) feed(arr string) int {
	for i := 0; i < d.capacity; i++ {
		d.eat(arr[i])
		if i+1 == d.capacity {
			return i
		}
	}
	return -1
}

func (d *pacman) eat(symbol byte) {
	if d.size == 0 && d.capacity > 0 {
		d.pushTail(symbol)
		d.diffElements++
		return
	}

	lastNum := d.tail.Symbol
	d.pushTail(symbol)
	if lastNum != symbol {
		d.diffElements++
	}

	if d.size > d.capacity {
		d.poop()
	}
}

func (d *pacman) poop() {
	val, ok := d.removeFront()
	if ok {
		if d.size == 0 {
			return
		}
		if d.head.Symbol == val {
			return
		}
		d.diffElements--
	}
}

func (d *pacman) removeFront() (byte, bool) {
	out := d.head
	if d.head != nil {
		d.head = d.head.Next
		if d.head == nil {
			d.tail = nil
		}
	}
	if out != nil {
		d.size--
		return out.Symbol, true
	}
	return 0, false
}

func (d *pacman) pushTail(symbol byte) {
	d.size++
	if d.head == nil && d.tail == nil {
		d.head = &Node{
			Symbol: symbol,
		}
		d.tail = d.head
		return
	}
	d.tail.Next = &Node{
		Symbol: symbol,
	}
	d.tail = d.tail.Next
}

type Node struct {
	Symbol byte
	Next   *Node
}
