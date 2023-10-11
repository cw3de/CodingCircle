package linkedlist

type Node struct {
	next  *Node
	value int
}

func NewNode(value int) *Node {
	return &Node{
		next:  nil,
		value: value,
	}
}
