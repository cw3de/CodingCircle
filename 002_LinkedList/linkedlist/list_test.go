package linkedlist_test

import (
	"testing"

	"github.com/cw3de/CodingCircle/LinkedList/linkedlist"
	"github.com/stretchr/testify/assert"
)

func TestBadArgument(t *testing.T) {
	l := linkedlist.New()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.RemoveFromBack(0)
	a := l.GetAsSlice()
	assert.Equal(t, a, []int{1, 2, 3, 4})
}

func TestEmpty(t *testing.T) {
	l := linkedlist.New()
	l.RemoveFromBack(2)
	a := l.GetAsSlice()
	assert.Equal(t, a, []int{})
}

func TestTooSmall(t *testing.T) {
	l := linkedlist.New()
	l.Append(1)
	l.RemoveFromBack(2)
	a := l.GetAsSlice()
	assert.Equal(t, a, []int{1})
}

func TestJustEnough(t *testing.T) {
	l := linkedlist.New()
	l.Append(1)
	l.Append(2)
	l.RemoveFromBack(2)
	a := l.GetAsSlice()
	assert.Equal(t, a, []int{2})
}

func TestMany(t *testing.T) {
	l := linkedlist.New()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Append(5)
	l.RemoveFromBack(2)
	a := l.GetAsSlice()
	assert.Equal(t, a, []int{1, 2, 3, 5})
}
