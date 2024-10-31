package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	extendOrigin(head)
	return divideLists(head)
}

func divideLists(head *Node) *Node {
	newHead := head.Next
	next := head.Next

	for {
		head.Next = next.Next
		head = head.Next
		if head == nil {
			next.Next = nil
			break
		}
		next.Next = next.Next.Next
		next = next.Next
	}
	return newHead
}

func extendOrigin(head *Node) {
	originHead := head

	for {
		if head == nil {
			break
		}
		copyItem := &Node{
			Val:  head.Val,
			Next: head.Next,
		}
		head.Next = copyItem
		head = head.Next.Next
	}

	head = originHead
	for {
		if head == nil {
			break
		}
		if head.Random == nil {
			head = head.Next.Next
			continue
		}
		head.Next.Random = head.Random.Next
		head = head.Next.Next
	}

	head = originHead
}
