package slist

type SList struct {
	Next  *SList
	Value int
}

func New(value int) *SList {
	return &SList{Value: value}
}

func (l *SList) CheckLoop() bool {
	slow := l
	fast := l
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
