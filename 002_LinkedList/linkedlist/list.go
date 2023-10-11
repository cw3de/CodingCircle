package linkedlist

import "fmt"

type List struct {
	head *Node
}

func New() *List {
	return &List{
		nil,
	}
}

func (l *List) Append(value int) {
	if l.head == nil {
		l.head = NewNode(value)
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = NewNode(value)
	}
}

func (l *List) GetAsSlice() []int {
	// var result []int

	result := make([]int, 0)

	c := l.head
	for c != nil {
		result = append(result, c.value)
		c = c.next
	}

	return result
}

func (l *List) RemoveFromBack(indexFromTail int) {
	if indexFromTail < 1 {
		fmt.Println(indexFromTail, "too small")
		return
	}

	// TODO: edge cases
	var previous *Node = l.head
	var current *Node = l.head
	// skip <indexFromTail> nodes
	for i := 0; i < indexFromTail; i++ {
		if current == nil {
			// fmt.Println("not enough elements")
			return
		}
		// fmt.Println("Skip", current.value)
		current = current.next
	}

	if current == nil {
		// liste hat genau <indexFromTail> Elemente
		l.head = l.head.next
	} else {
		current = current.next

		for current != nil {
			// fmt.Println("Skip", current.value, "and", previous.value)
			previous = previous.next
			current = current.next
		}
		previous.next = previous.next.next
	}
}
