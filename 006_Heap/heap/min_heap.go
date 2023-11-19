package heap

import "errors"

type Item[TValue any] struct {
	Value    TValue
	Priority int
}
type MinHeap[TValue any] struct {
	items []Item[TValue]
}

func New[TValue any]() *MinHeap[TValue] {
	return &MinHeap[TValue]{}
}

func leftChildIndex(parentIndex int) int {
	return 2*parentIndex + 1
}

func rightChildIndex(parentIndex int) int {
	return 2*parentIndex + 2
}

func parentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

func (h *MinHeap[TValue]) Items() []Item[TValue] {
	return h.items
}

func (h *MinHeap[TValue]) Peek() *Item[TValue] {
	if len(h.items) == 0 {
		return nil
	}

	return &h.items[0]
}

func (h *MinHeap[TValue]) isValid() {
	for i := 1; i < len(h.items); i++ {
		if h.items[i].Priority < h.items[parentIndex(i)].Priority {
			panic("Invalid heap")
		}
	}
}

func (h *MinHeap[TValue]) Insert(value TValue, priority int) {
	defer h.isValid()

	h.items = append(h.items, Item[TValue]{value, priority})

	h.siftUp(len(h.items) - 1)
}

func (h *MinHeap[TValue]) siftUp(index int) {

	for index > 0 {
		parent := parentIndex(index)

		if h.items[index].Priority > h.items[parent].Priority {
			return
		}
		h.items[index], h.items[parent] = h.items[parent], h.items[index]
		index = parent
	}
}

func (h *MinHeap[TValue]) Remove(index int) error {
	if index < 0 || index >= len(h.items) {
		return errors.New("index out of range")
	}

	defer h.isValid()

	h.items[index] = h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]

	h.siftDown(index)

	return nil
}

func (h *MinHeap[TValue]) siftDown(index int) {
	left := leftChildIndex(index)
	right := rightChildIndex(index)
	smallest := index

	if left < len(h.items) && h.items[left].Priority < h.items[smallest].Priority {
		smallest = left
	}

	if right < len(h.items) && h.items[right].Priority < h.items[smallest].Priority {
		smallest = right
	}

	if index != smallest {
		h.items[index], h.items[smallest] = h.items[smallest], h.items[index]
		h.siftDown(smallest)
	}
}

func (h *MinHeap[TValue]) Update(index int, newPriority int) error {
	if index < 0 || index >= len(h.items) {
		return errors.New("index out of range")
	}

	defer h.isValid()

	oldPriority := h.items[index].Priority
	h.items[index].Priority = newPriority

	if newPriority < oldPriority {
		h.siftUp(index)
	} else {
		h.siftDown(index)
	}
	return nil
}
